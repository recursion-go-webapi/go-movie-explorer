import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Layout } from '@/components/Layout';
import { HomePage } from '@/pages/HomePage';
import { MoviesPage } from '@/pages/MoviesPage';
import { MovieDetailPage } from '@/pages/MovieDetailPage';
import { GenrePage } from '@/pages/GenrePage';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<HomePage />} />
          <Route path="movies" element={<MoviesPage />} />
          <Route path="genre" element={<GenrePage />} />
          <Route path="movie/:id" element={<MovieDetailPage />} />
          <Route path="genre/:id" element={<GenrePage />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App
