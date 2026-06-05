import { useState, useEffect, useRef, useCallback } from "react";
import type { ReactElement } from "react";

const SIEVE_LIMIT = 100_000_000;

function buildSieve(n: number): Uint8Array {
  const arr = new Uint8Array(n + 1).fill(1);
  arr[0] = 0;
  arr[1] = 0;
  for (let i = 2; i * i <= n; i++) {
    if (arr[i]) {
      for (let j = i * i; j <= n; j += i) arr[j] = 0;
    }
  }
  return arr;
}

let _sieve: Uint8Array | null = null;
function getSieve(): Uint8Array {
  if (!_sieve) _sieve = buildSieve(SIEVE_LIMIT);
  return _sieve;
}

function isPrime(n: number): boolean {
  if (n <= SIEVE_LIMIT) return getSieve()[n] === 1;
  if (n % 2 === 0) return false;
  for (let i = 3; i <= Math.sqrt(n); i += 2) if (n % i === 0) return false;
  return true;
}

function gen67Numbers(): number[] {
  const result: number[] = [];
  const queue: string[] = ["6", "7"];

  while (queue.length > 0) {
    const s = queue.shift();
    if (!s) break;

    const n = parseInt(s, 10);
    
    if (n <= SIEVE_LIMIT) {
      result.push(n);
      
      const next6 = s + "6";
      const next7 = s + "7";
      
      if (parseInt(next6, 10) <= SIEVE_LIMIT) {
        queue.push(next6);
      }
      if (parseInt(next7, 10) <= SIEVE_LIMIT) {
        queue.push(next7);
      }
    }
  }

  return result.sort((a, b) => a - b);
}

const NUMBERS_67 = gen67Numbers();

interface Level {
  name: string;
  threshold: number;
  nextAt: number;
  color: string;
}

const LEVELS: Level[] = [
  { name: "Iniciante",     threshold: 0,    nextAt: 20,       color: "#9CA3AF" },
  { name: "Aprendiz",       threshold: 20,   nextAt: 60,       color: "#60A5FA" },
  { name: "Médio",          threshold: 60,   nextAt: 150,      color: "#34D399" },
  { name: "Avançado",       threshold: 150,  nextAt: 350,      color: "#FBBF24" },
  { name: "Mestre",         threshold: 350,  nextAt: 700,      color: "#F97316" },
  { name: "Lendário",       threshold: 700,  nextAt: 1200,     color: "#A78BFA" },
  { name: "Transcendente",  threshold: 1200, nextAt: Infinity, color: "#F472B6" },
];

function getLevel(aura: number): Level {
  for (let i = LEVELS.length - 1; i >= 0; i--) {
    if (aura >= LEVELS[i].threshold) return LEVELS[i];
  }
  return LEVELS[0];
}

interface LevelProgress {
  pct: number;
  current: number;
  total: number;
  isMax: boolean;
}

function getLevelProgress(aura: number): LevelProgress {
  const lvl = getLevel(aura);
  const next = lvl.nextAt === Infinity ? lvl.threshold + 500 : lvl.nextAt;
  const pct = Math.min(100, ((aura - lvl.threshold) / (next - lvl.threshold)) * 100);
  return { pct, current: aura - lvl.threshold, total: next - lvl.threshold, isMax: lvl.nextAt === Infinity };
}

// ─── Partículas de aura ───────────────────────────────────────────────────────
interface ParticlesProps {
  active: boolean;
  color: string;
}

function Particles({ active, color }: ParticlesProps): ReactElement | null {
  const particles = Array.from({ length: 12 }, (_, i) => i);
  if (!active) return null;
  return (
    <div style={{ position: "absolute", inset: 0, pointerEvents: "none", overflow: "hidden" }}>
      {particles.map((i) => {
        const angle = (i / 12) * 360;
        const dist = 40 + Math.random() * 60;
        const tx = Math.cos((angle * Math.PI) / 180) * dist;
        const ty = Math.sin((angle * Math.PI) / 180) * dist;
        return (
          <div
            key={i}
            style={{
              position: "absolute",
              left: "50%",
              top: "50%",
              width: 8,
              height: 8,
              borderRadius: "50%",
              background: color,
              transform: "translate(-50%, -50%)",
              animation: `particle-burst 0.7s ease-out forwards`,
              animationDelay: `${i * 20}ms`,
              "--tx": `${tx}px`,
              "--ty": `${ty}px`,
            } as React.CSSProperties & Record<string, string | number>}
          />
        );
      })}
    </div>
  );
}

interface Feedback {
  correct: boolean;
  message: string;
  gain: number;
}

interface FeedbackImg {
  src: string;
  visible: boolean;
}

interface AuraFloat {
  id: number;
  text: string;
  color: string;
  x: number;
}

export default function App(): ReactElement {
  const [sieveReady, setSieveReady] = useState<boolean>(false);
  const [current, setCurrent] = useState<number | null>(null);
  const [aura, setAura] = useState<number>(0);
  const [hits, setHits] = useState<number>(0);
  const [misses, setMisses] = useState<number>(0);
  const [streak, setStreak] = useState<number>(0);
  const [bestStreak, setBestStreak] = useState<number>(0);
  const [waiting, setWaiting] = useState<boolean>(false);
  const [feedback, setFeedback] = useState<Feedback | null>(null);
  const [feedbackImg, setFeedbackImg] = useState<FeedbackImg | null>(null);
  const [popAnim, setPopAnim] = useState<boolean>(false);
  const [particles, setParticles] = useState<boolean>(false);
  const [shakeAnim, setShakeAnim] = useState<boolean>(false);
  const [auraFloats, setAuraFloats] = useState<AuraFloat[]>([]);

  const imgFadeTimer = useRef<ReturnType<typeof setTimeout> | null>(null);
  const feedbackTimer = useRef<ReturnType<typeof setTimeout> | null>(null);
  const particleTimer = useRef<ReturnType<typeof setTimeout> | null>(null);
  const floatIdRef = useRef<number>(0);
  const [_currentIndex, setCurrentIndex] = useState<number>(0);

  const pickNext = useCallback((): void => {
    setCurrentIndex((prevIndex) => {
      const nextIndex = prevIndex + 1;
      // Se passar do tamanho da lista, reinicia do 0
      const safeIndex = nextIndex < NUMBERS_67.length ? nextIndex : 0;
      
      setCurrent(NUMBERS_67[safeIndex]);
      return safeIndex;
    });
  }, []);

  useEffect(() => {
    const t = setTimeout(() => {
      getSieve();
      setSieveReady(true);
      
      setCurrent(NUMBERS_67[0]);
      setCurrentIndex(0);
    }, 50);
    
    return () => clearTimeout(t);
  }, []);

  const spawnFloat = useCallback((text: string, color: string): void => {
    const id = ++floatIdRef.current;
    const x = 35 + Math.random() * 30;
    setAuraFloats((prev) => [...prev, { id, text, color, x }]);
    setTimeout(() => {
      setAuraFloats((prev) => prev.filter((f) => f.id !== id));
    }, 1200);
  }, []);

  const clearImgTimers = (): void => {
    if (imgFadeTimer.current) clearTimeout(imgFadeTimer.current);
  };

  const answer = useCallback(
    (userSaysIsPrime: boolean): void => {
      if (waiting || current === null || !sieveReady) return;
      setWaiting(true);

      const correct = isPrime(current);
      const right = userSaysIsPrime === correct;
      const bonusMult = streak >= 5 ? 3 : streak >= 3 ? 2 : 1;
      const gain = right ? 10 * bonusMult : 0;

      if (right) {
        const newAura = aura + gain;
        const newStreak = streak + 1;
        const newHits = hits + 1;
        setAura(newAura);
        setStreak(newStreak);
        setHits(newHits);
        setBestStreak((b) => Math.max(b, newStreak));
        setFeedback({
          correct: true,
          message: correct
            ? `${current.toLocaleString("pt-BR")} é primo!`
            : `${current.toLocaleString("pt-BR")} não é primo!`,
          gain,
        });
        setPopAnim(true);
        setParticles(true);
        spawnFloat(`+${gain} ✨`, "#34D399");
        if (particleTimer.current) clearTimeout(particleTimer.current);
        particleTimer.current = setTimeout(() => setParticles(false), 800);
        setTimeout(() => setPopAnim(false), 300);

        // Imagem de acerto
        const idx = Math.floor(Math.random() * 4) + 1;
        clearImgTimers();
        setFeedbackImg({ src: `aura${idx}.png`, visible: true });
        imgFadeTimer.current = setTimeout(() => {
          setFeedbackImg((fi) => fi ? { ...fi, visible: false } : fi);
          imgFadeTimer.current = setTimeout(() => setFeedbackImg(null), 2600);
        }, 1800);
      } else {
        setMisses((m) => m + 1);
        setStreak(0);
        setFeedback({
          correct: false,
          message: correct
            ? `Errou! ${current.toLocaleString("pt-BR")} é primo.`
            : `Errou! ${current.toLocaleString("pt-BR")} não é primo.`,
          gain: 0,
        });
        setShakeAnim(true);
        spawnFloat("-♻️", "#F87171");
        setTimeout(() => setShakeAnim(false), 500);

        const idx = Math.floor(Math.random() * 2) + 1;
        clearImgTimers();
        setFeedbackImg({ src: `beta${idx}.png`, visible: true });
        imgFadeTimer.current = setTimeout(() => {
          setFeedbackImg((fi) => fi ? { ...fi, visible: false } : fi);
          imgFadeTimer.current = setTimeout(() => setFeedbackImg(null), 2600);
        }, 1800);
      }

      if (feedbackTimer.current) clearTimeout(feedbackTimer.current);
      feedbackTimer.current = setTimeout(() => {
        setFeedback(null);
        setWaiting(false);
        pickNext();
      }, 2200);
    },
    [waiting, current, sieveReady, aura, hits, streak, pickNext, spawnFloat]
  );

  const lvl = getLevel(aura);
  const progress = getLevelProgress(aura);
  const accuracy = hits + misses > 0 ? Math.round((hits / (hits + misses)) * 100) : 0;

  return (
    <>
      <style>{`
        @import url('https://fonts.googleapis.com/css2?family=Rajdhani:wght@400;500;600;700&family=JetBrains+Mono:wght@400;700&display=swap');

        * { box-sizing: border-box; margin: 0; padding: 0; }

        body {
          background: #080B14;
          min-height: 100vh;
        }

        .game-root {
          background: #080B14;
          min-height: 100vh;
          display: flex;
          flex-direction: column;
          align-items: center;
          padding: 2rem 1rem;
          font-family: 'Rajdhani', sans-serif;
          color: #E2E8F0;
          position: relative;
          overflow: hidden;
        }

        .bg-grid {
          position: fixed;
          inset: 0;
          background-image:
            linear-gradient(rgba(52, 211, 153, 0.03) 1px, transparent 1px),
            linear-gradient(90deg, rgba(52, 211, 153, 0.03) 1px, transparent 1px);
          background-size: 40px 40px;
          pointer-events: none;
          z-index: 0;
        }

        .bg-glow {
          position: fixed;
          width: 600px;
          height: 600px;
          border-radius: 50%;
          background: radial-gradient(circle, rgba(52,211,153,0.07) 0%, transparent 70%);
          top: -200px;
          left: 50%;
          transform: translateX(-50%);
          pointer-events: none;
          z-index: 0;
        }

        .content {
          position: relative;
          z-index: 1;
          width: 100%;
          max-width: 520px;
          display: flex;
          flex-direction: column;
          align-items: center;
          gap: 1.5rem;
        }

        /* Header */
        .header {
          text-align: center;
        }

        .title {
          font-size: 13px;
          letter-spacing: 0.25em;
          text-transform: uppercase;
          color: #34D399;
          margin-bottom: 0.25rem;
        }

        .subtitle {
          font-size: 11px;
          letter-spacing: 0.15em;
          text-transform: uppercase;
          color: #4B5563;
        }

        /* Stats row */
        .stats-row {
          display: grid;
          grid-template-columns: repeat(4, 1fr);
          gap: 8px;
          width: 100%;
        }

        .stat-card {
          background: #0F1420;
          border: 1px solid #1E293B;
          border-radius: 8px;
          padding: 0.75rem 0.5rem;
          text-align: center;
        }

        .stat-label {
          font-size: 10px;
          letter-spacing: 0.12em;
          text-transform: uppercase;
          color: #4B5563;
          margin-bottom: 4px;
        }

        .stat-val {
          font-family: 'JetBrains Mono', monospace;
          font-size: 20px;
          font-weight: 700;
          color: #E2E8F0;
        }

        .stat-val.aura { color: #34D399; }
        .stat-val.streak { color: #FBBF24; }
        .stat-val.accuracy { color: #60A5FA; }

        /* Level bar */
        .level-section {
          width: 100%;
          background: #0F1420;
          border: 1px solid #1E293B;
          border-radius: 8px;
          padding: 0.875rem 1rem;
        }

        .level-header {
          display: flex;
          justify-content: space-between;
          align-items: center;
          margin-bottom: 8px;
        }

        .level-name {
          font-size: 11px;
          letter-spacing: 0.15em;
          text-transform: uppercase;
          font-weight: 600;
        }

        .level-progress-text {
          font-family: 'JetBrains Mono', monospace;
          font-size: 11px;
          color: #4B5563;
        }

        .level-bar-bg {
          height: 4px;
          background: #1E293B;
          border-radius: 99px;
          overflow: hidden;
        }

        .level-bar-fill {
          height: 100%;
          border-radius: 99px;
          transition: width 0.5s cubic-bezier(0.4, 2, 0.6, 1);
        }

        /* Number arena */
        .arena {
          width: 100%;
          background: #0F1420;
          border: 1px solid #1E293B;
          border-radius: 12px;
          padding: 2rem 1.5rem;
          text-align: center;
          position: relative;
          overflow: hidden;
          min-height: 160px;
          display: flex;
          flex-direction: column;
          align-items: center;
          justify-content: center;
        }

        .arena-scanline {
          position: absolute;
          inset: 0;
          background: repeating-linear-gradient(
            0deg,
            transparent,
            transparent 2px,
            rgba(52, 211, 153, 0.015) 2px,
            rgba(52, 211, 153, 0.015) 4px
          );
          pointer-events: none;
        }

        .number-label {
          font-size: 10px;
          letter-spacing: 0.2em;
          text-transform: uppercase;
          color: #4B5563;
          margin-bottom: 0.5rem;
        }

        .number-display {
          font-family: 'JetBrains Mono', monospace;
          font-size: 52px;
          font-weight: 700;
          color: #E2E8F0;
          line-height: 1;
          transition: transform 0.15s;
          letter-spacing: -2px;
        }

        .number-display.pop {
          animation: pop 0.3s ease;
        }

        .number-display.shake {
          animation: shake 0.5s ease;
        }

        .number-sub {
          font-size: 11px;
          letter-spacing: 0.2em;
          text-transform: uppercase;
          color: #374151;
          margin-top: 0.5rem;
        }

        .float-container {
          position: absolute;
          inset: 0;
          pointer-events: none;
          overflow: hidden;
        }

        .float-text {
          position: absolute;
          bottom: 40%;
          font-family: 'JetBrains Mono', monospace;
          font-size: 18px;
          font-weight: 700;
          animation: float-up 1.2s ease-out forwards;
          pointer-events: none;
        }

        /* Feedback */
        .feedback-area {
          width: 100%;
          min-height: 64px;
          display: flex;
          flex-direction: column;
          align-items: center;
          gap: 0.5rem;
          position: relative;
        }

        .feedback-img {
          width: 100px;
          height: 100px;
          object-fit: contain;
          border-radius: 12px;
          transition: opacity 2.5s ease;
        }

        .feedback-msg {
          font-size: 14px;
          font-weight: 600;
          letter-spacing: 0.05em;
          text-align: center;
          animation: fade-in-msg 0.2s ease;
        }

        .feedback-msg.correct { color: #34D399; }
        .feedback-msg.wrong   { color: #F87171; }

        .gain-badge {
          font-family: 'JetBrains Mono', monospace;
          font-size: 12px;
          padding: 2px 10px;
          border-radius: 99px;
          background: rgba(52, 211, 153, 0.12);
          color: #34D399;
          border: 1px solid rgba(52, 211, 153, 0.25);
        }

        /* Streak badge */
        .streak-badge {
          font-size: 11px;
          letter-spacing: 0.1em;
          padding: 4px 12px;
          border-radius: 99px;
          font-weight: 600;
          transition: all 0.3s;
        }

        .streak-badge.hot {
          background: rgba(251, 191, 36, 0.12);
          color: #FBBF24;
          border: 1px solid rgba(251, 191, 36, 0.25);
        }

        .streak-badge.warm {
          background: rgba(249, 115, 22, 0.1);
          color: #F97316;
          border: 1px solid rgba(249, 115, 22, 0.2);
        }

        /* Buttons */
        .btn-row {
          display: grid;
          grid-template-columns: 1fr 1fr;
          gap: 12px;
          width: 100%;
        }

        .btn-prime, .btn-not {
          padding: 14px;
          font-family: 'Rajdhani', sans-serif;
          font-size: 15px;
          font-weight: 700;
          letter-spacing: 0.12em;
          text-transform: uppercase;
          border-radius: 8px;
          cursor: pointer;
          border: 1px solid;
          transition: all 0.15s;
          position: relative;
          overflow: hidden;
        }

        .btn-prime:not(:disabled):hover,
        .btn-not:not(:disabled):hover {
          transform: translateY(-1px);
        }

        .btn-prime:not(:disabled):active,
        .btn-not:not(:disabled):active {
          transform: translateY(1px) scale(0.98);
        }

        .btn-prime {
          background: rgba(52, 211, 153, 0.08);
          color: #34D399;
          border-color: rgba(52, 211, 153, 0.3);
        }

        .btn-prime:not(:disabled):hover {
          background: rgba(52, 211, 153, 0.15);
          border-color: rgba(52, 211, 153, 0.5);
          box-shadow: 0 0 20px rgba(52, 211, 153, 0.15);
        }

        .btn-not {
          background: rgba(248, 113, 113, 0.08);
          color: #F87171;
          border-color: rgba(248, 113, 113, 0.3);
        }

        .btn-not:not(:disabled):hover {
          background: rgba(248, 113, 113, 0.15);
          border-color: rgba(248, 113, 113, 0.5);
          box-shadow: 0 0 20px rgba(248, 113, 113, 0.15);
        }

        .btn-prime:disabled,
        .btn-not:disabled {
          opacity: 0.3;
          cursor: not-allowed;
        }

        /* Hint */
        .hint {
          font-size: 11px;
          color: #374151;
          letter-spacing: 0.08em;
          text-align: center;
        }

        .hint span {
          color: #4B5563;
        }

        /* Loading */
        .loading {
          display: flex;
          flex-direction: column;
          align-items: center;
          gap: 1rem;
          padding: 3rem;
          color: #4B5563;
          font-size: 13px;
          letter-spacing: 0.15em;
          text-transform: uppercase;
        }

        .loading-dot {
          width: 8px;
          height: 8px;
          border-radius: 50%;
          background: #34D399;
          animation: pulse-dot 1s ease-in-out infinite;
        }

        /* Animations */
        @keyframes pop {
          0%   { transform: scale(1); }
          40%  { transform: scale(1.18); }
          100% { transform: scale(1); }
        }

        @keyframes shake {
          0%   { transform: translateX(0); }
          20%  { transform: translateX(-8px); }
          40%  { transform: translateX(8px); }
          60%  { transform: translateX(-5px); }
          80%  { transform: translateX(5px); }
          100% { transform: translateX(0); }
        }

        @keyframes float-up {
          0%   { transform: translateY(0); opacity: 1; }
          100% { transform: translateY(-80px); opacity: 0; }
        }

        @keyframes fade-in-msg {
          from { opacity: 0; transform: translateY(4px); }
          to   { opacity: 1; transform: translateY(0); }
        }

        @keyframes pulse-dot {
          0%, 100% { opacity: 0.3; transform: scale(0.8); }
          50%       { opacity: 1; transform: scale(1.2); }
        }

        @keyframes particle-burst {
          0%   { transform: translate(-50%, -50%) scale(1); opacity: 1; }
          100% { transform: translate(calc(-50% + var(--tx)), calc(-50% + var(--ty))) scale(0); opacity: 0; }
        }

        @keyframes slide-in {
          from { opacity: 0; transform: translateY(16px); }
          to   { opacity: 1; transform: translateY(0); }
        }

        .content { animation: slide-in 0.4s ease; }
      `}</style>

      <div className="game-root">
        <div className="bg-grid" />
        <div className="bg-glow" />

        {!sieveReady ? (
          <div className="loading">
            <div className="loading-dot" />
            Construindo o crivo…
          </div>
        ) : (
          <div className="content">
            {/* Header */}
            <div className="header">
              <div className="title">✨ Prime Aura</div>
              <div className="subtitle">números com algarismos 6 e 7</div>
            </div>

            {/* Stats */}
            <div className="stats-row">
              <div className="stat-card">
                <div className="stat-label">Aura</div>
                <div className="stat-val aura">{aura}</div>
              </div>
              <div className="stat-card">
                <div className="stat-label">Acertos</div>
                <div className="stat-val">{hits}</div>
              </div>
              <div className="stat-card">
                <div className="stat-label">Erros</div>
                <div className="stat-val">{misses}</div>
              </div>
              <div className="stat-card">
                <div className="stat-label">Sequência</div>
                <div className="stat-val streak">{streak}</div>
              </div>
            </div>

            {/* Level bar */}
            <div className="level-section">
              <div className="level-header">
                <span className="level-name" style={{ color: lvl.color }}>
                  {lvl.name}
                </span>
                {!progress.isMax ? (
                  <span className="level-progress-text">
                    {progress.current} / {progress.total}
                  </span>
                ) : (
                  <span className="level-progress-text" style={{ color: lvl.color }}>
                    nível máximo
                  </span>
                )}
              </div>
              <div className="level-bar-bg">
                <div
                  className="level-bar-fill"
                  style={{ width: `${progress.pct}%`, background: lvl.color }}
                />
              </div>
            </div>

            {/* Streak badge */}
            {streak >= 3 && (
              <div className={`streak-badge ${streak >= 5 ? "hot" : "warm"}`}>
                {streak >= 5 ? `🔥 sequência ×${streak}! bônus 3×` : `⚡ sequência ×${streak} — bônus 2×`}
              </div>
            )}

            {/* Arena */}
            <div className="arena">
              <div className="arena-scanline" />

              <Particles active={particles} color={lvl.color} />

              <div className="float-container">
                {auraFloats.map((f) => (
                  <div
                    key={f.id}
                    className="float-text"
                    style={{ left: `${f.x}%`, color: f.color }}
                  >
                    {f.text}
                  </div>
                ))}
              </div>

              <div className="number-label">é primo?</div>
              <div
                className={`number-display ${popAnim ? "pop" : ""} ${shakeAnim ? "shake" : ""}`}
              >
                {current !== null ? current.toLocaleString("pt-BR") : "—"}
              </div>
              <div className="number-sub">algarismos 6 e 7</div>
            </div>

            {/* Feedback */}
            <div className="feedback-area">
              {feedbackImg && (
                <img
                  className="feedback-img"
                  src={feedbackImg.src}
                  alt=""
                  style={{ opacity: feedbackImg.visible ? 1 : 0 }}
                />
              )}
              {feedback && (
                <>
                  <div className={`feedback-msg ${feedback.correct ? "correct" : "wrong"}`}>
                    {feedback.message}
                  </div>
                  {feedback.gain > 0 && (
                    <div className="gain-badge">+{feedback.gain} aura</div>
                  )}
                </>
              )}
            </div>

            {/* Buttons */}
            <div className="btn-row">
              <button
                className="btn-prime"
                onClick={() => answer(true)}
                disabled={waiting || !sieveReady}
              >
                ✅ É primo
              </button>
              <button
                className="btn-not"
                onClick={() => answer(false)}
                disabled={waiting || !sieveReady}
              >
                ❌ Não é primo
              </button>
            </div>

            {/* Hint */}
            <div className="hint">
              melhor sequência: <span>{bestStreak}</span> &nbsp;·&nbsp; precisão:{" "}
              <span>{accuracy}%</span>
            </div>
          </div>
        )}
      </div>
    </>
  );
}