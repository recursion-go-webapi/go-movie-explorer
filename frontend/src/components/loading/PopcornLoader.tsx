import { motion } from "framer-motion";

interface PopcornLoaderProps {
  size?: "sm" | "md" | "lg";
  className?: string;
}

export function PopcornLoader({ size = "md", className = "" }: PopcornLoaderProps) {
  const sizeClasses = {
    sm: "w-6 h-8",
    md: "w-8 h-10",
    lg: "w-10 h-12"
  };

  return (
    <div className={`flex items-center justify-center ${className}`}>
      <div className={`${sizeClasses[size]} relative`}>
        <div className="absolute bottom-0 left-0 right-0 h-2/3 bg-gradient-to-t from-red-600 to-red-500 rounded-b-lg border-2 border-amber-400" />
        
        <div className="absolute bottom-1/3 left-1/2 -translate-x-1/2 w-1/3 h-1/6 bg-yellow-300 rounded-full opacity-80" />
        
        {Array.from({ length: 4 }).map((_, i) => (
          <motion.div
            key={i}
            animate={{
              y: [0, -8, 0],
              rotate: [0, 15, -15, 0],
              scale: [1, 1.1, 1],
            }}
            transition={{
              duration: 2,
              repeat: Infinity,
              delay: i * 0.3,
              ease: "easeInOut",
            }}
            className="absolute w-1 h-1 bg-yellow-200 rounded-full"
            style={{
              top: `${20 + i * 5}%`,
              left: `${30 + (i % 2) * 40}%`,
            }}
          />
        ))}
        
        {Array.from({ length: 3 }).map((_, i) => (
          <motion.div
            key={i}
            animate={{
              y: [0, -6, 0],
              x: [0, 2, -2, 0],
              scale: [1, 1.2, 1],
            }}
            transition={{
              duration: 2.5,
              repeat: Infinity,
              delay: i * 0.4,
              ease: "easeInOut",
            }}
            className="absolute w-1 h-1 bg-yellow-100 rounded-full"
            style={{
              top: `${10 + i * 8}%`,
              left: `${40 + (i % 3) * 20}%`,
            }}
          />
        ))}
      </div>
    </div>
  );
}