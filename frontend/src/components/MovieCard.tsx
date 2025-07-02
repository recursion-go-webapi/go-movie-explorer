import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import type { Movie } from "@/types/movie";
import { getImageUrl, formatDate, formatRating, truncateText } from "@/lib/utils";

interface MovieCardProps {
  movie: Movie;
  onClick?: (movie: Movie) => void;
}

export function MovieCard({ movie, onClick }: MovieCardProps) {
  return (
    <Card 
      className="cursor-pointer transition-transform hover:scale-105 hover:shadow-lg"
      onClick={() => onClick?.(movie)}
    >
      <CardHeader className="p-0">
        <div className="aspect-[2/3] overflow-hidden rounded-t-lg bg-gray-200 flex items-center justify-center">
          {movie.poster_path ? (
            <img
              src={getImageUrl(movie.poster_path, 'w500')}
              alt={movie.title}
              className="w-full h-full object-cover"
              loading="lazy"
              onError={(e) => {
                const parent = e.currentTarget.parentElement;
                if (parent) {
                  parent.innerHTML = '<div class="text-gray-500 text-4xl">🎬</div>';
                }
              }}
            />
          ) : (
            <div className="text-gray-500 text-4xl">🎬</div>
          )}
        </div>
      </CardHeader>
      <CardContent className="p-4">
        <CardTitle className="text-sm font-semibold mb-2 line-clamp-2">
          {movie.title}
        </CardTitle>
        <CardDescription className="text-xs mb-2">
          {formatDate(movie.release_date)}
        </CardDescription>
        <div className="flex items-center justify-between text-xs">
          <span className="flex items-center">
            ⭐ {formatRating(movie.vote_average)}
          </span>
          <span className="text-muted-foreground">
            人気度: {Math.round(movie.popularity)}
          </span>
        </div>
        {movie.overview && (
          <p className="text-xs text-muted-foreground mt-2 line-clamp-3">
            {truncateText(movie.overview, 100)}
          </p>
        )}
      </CardContent>
    </Card>
  );
}