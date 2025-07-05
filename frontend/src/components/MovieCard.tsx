import { motion } from "framer-motion";
import type { Movie } from "@/types/movie";
import { getImageUrl, formatDate, formatRating, truncateText } from "@/lib/utils";

interface MovieCardProps {
  movie: Movie;
  onClick?: (movie: Movie) => void;
}

export function MovieCard({ movie, onClick }: MovieCardProps) {
  return (
    <motion.div
      whileHover={{ 
        scale: 1.05, 
        y: -10,
        transition: { duration: 0.3, ease: "easeOut" }
      }}
      whileTap={{ scale: 0.95 }}
      className="group cursor-pointer relative overflow-hidden rounded-xl bg-gradient-to-br from-slate-900 to-slate-800 shadow-2xl border border-slate-700/50"
      onClick={() => onClick?.(movie)}
    >
      <div className="aspect-[2/3] overflow-hidden relative">
        <motion.div
          whileHover={{ scale: 1.1 }}
          transition={{ duration: 0.5, ease: "easeOut" }}
          className="w-full h-full"
        >
          {movie.poster_path ? (
            <img
              src={getImageUrl(movie.poster_path, 'w500')}
              alt={movie.title}
              className="w-full h-full object-cover"
              loading="lazy"
              onError={(e) => {
                const parent = e.currentTarget.parentElement;
                if (parent) {
                  parent.innerHTML = '<div class="flex items-center justify-center w-full h-full text-amber-400 text-6xl">🎬</div>';
                }
              }}
            />
          ) : (
            <div className="flex items-center justify-center w-full h-full text-amber-400 text-6xl">🎬</div>
          )}
        </motion.div>
        
        <div className="absolute inset-0 bg-gradient-to-t from-black/80 via-transparent to-transparent opacity-0 group-hover:opacity-100 transition-opacity duration-300" />
        
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          whileHover={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.3, ease: "easeOut" }}
          className="absolute bottom-0 left-0 right-0 p-4 text-white opacity-0 group-hover:opacity-100"
        >
          <div className="flex items-center justify-between mb-2">
            <div className="flex items-center space-x-1">
              <span className="text-amber-400 text-lg">★</span>
              <span className="text-sm font-semibold">{formatRating(movie.vote_average)}</span>
            </div>
            <div className="text-xs text-slate-300">
              人気度: {Math.round(movie.popularity)}
            </div>
          </div>
          {movie.overview && (
            <p className="text-xs text-slate-200 leading-relaxed line-clamp-3">
              {truncateText(movie.overview, 120)}
            </p>
          )}
        </motion.div>
      </div>
      
      <div className="p-4 bg-gradient-to-br from-slate-800 to-slate-900">
        <h3 className="text-white font-bold text-sm mb-1 line-clamp-2 leading-tight">
          {movie.title}
        </h3>
        <p className="text-amber-400 text-xs font-medium">
          {formatDate(movie.release_date)}
        </p>
      </div>
      
      <div className="absolute top-2 right-2 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
        <div className="w-8 h-8 bg-amber-500 rounded-full flex items-center justify-center shadow-lg">
          <span className="text-slate-900 text-xs font-bold">▶</span>
        </div>
      </div>
    </motion.div>
  );
}