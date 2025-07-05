import { motion } from "framer-motion";

interface CameraLoaderProps {
  size?: "sm" | "md" | "lg";
  className?: string;
}

export function CameraLoader({ size = "md", className = "" }: CameraLoaderProps) {
  const sizeClasses = {
    sm: "w-8 h-6",
    md: "w-12 h-9",
    lg: "w-16 h-12"
  };

  return (
    <div className={`flex items-center justify-center ${className}`}>
      <div className={`${sizeClasses[size]} relative`}>
        <div className="absolute inset-0 bg-slate-700 rounded-lg border-2 border-amber-400">
          <div className="absolute top-1/2 left-1/2 w-1/3 h-1/3 bg-slate-800 rounded-full -translate-x-1/2 -translate-y-1/2 border border-amber-400">
            <motion.div
              animate={{ scale: [1, 1.2, 1] }}
              transition={{ duration: 1.5, repeat: Infinity, ease: "easeInOut" }}
              className="absolute inset-1 bg-amber-400 rounded-full opacity-60"
            />
          </div>
          
          <div className="absolute -top-1 left-1/4 w-1/2 h-1 bg-slate-600 rounded-t-sm" />
          
          <motion.div
            animate={{ opacity: [1, 0.3, 1] }}
            transition={{ duration: 2, repeat: Infinity, ease: "easeInOut" }}
            className="absolute top-1 right-1 w-1 h-1 bg-red-400 rounded-full"
          />
        </div>
        
        <motion.div
          animate={{ x: [0, 2, 0] }}
          transition={{ duration: 3, repeat: Infinity, ease: "easeInOut" }}
          className="absolute -right-1 top-1/2 w-1 h-1/2 bg-amber-400 rounded-r-sm -translate-y-1/2"
        />
      </div>
    </div>
  );
}