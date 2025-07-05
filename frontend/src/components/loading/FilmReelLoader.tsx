import { motion } from "framer-motion";

interface FilmReelLoaderProps {
  size?: "sm" | "md" | "lg";
  className?: string;
}

export function FilmReelLoader({ size = "md", className = "" }: FilmReelLoaderProps) {
  const sizeClasses = {
    sm: "w-8 h-8",
    md: "w-12 h-12",
    lg: "w-16 h-16"
  };

  return (
    <div className={`flex items-center justify-center ${className}`}>
      <motion.div
        animate={{ rotate: 360 }}
        transition={{ duration: 2, repeat: Infinity, ease: "linear" }}
        className={`${sizeClasses[size]} relative`}
      >
        <div className="absolute inset-0 rounded-full border-4 border-amber-400 bg-slate-800">
          <div className="absolute inset-2 rounded-full border-2 border-amber-400">
            <div className="absolute inset-1 rounded-full bg-slate-700">
              <div className="absolute top-1/2 left-1/2 w-1 h-1 bg-amber-400 rounded-full -translate-x-1/2 -translate-y-1/2" />
            </div>
          </div>
        </div>
        
        {Array.from({ length: 8 }).map((_, i) => (
          <div
            key={i}
            className="absolute w-1 h-1 bg-amber-400 rounded-full"
            style={{
              top: "50%",
              left: "50%",
              transform: `translate(-50%, -50%) rotate(${i * 45}deg) translateY(-${size === "sm" ? "12" : size === "md" ? "18" : "24"}px)`,
            }}
          />
        ))}
      </motion.div>
    </div>
  );
}