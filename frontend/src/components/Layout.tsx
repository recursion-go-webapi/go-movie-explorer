import { Outlet } from 'react-router-dom';
import { Navigation } from './Navigation';

export function Layout() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-indigo-500 via-purple-500 to-pink-500">
      <Navigation />
      <main className="pt-20 pb-8">
        <div className="container mx-auto px-4">
          <Outlet />
        </div>
      </main>
    </div>
  );
}