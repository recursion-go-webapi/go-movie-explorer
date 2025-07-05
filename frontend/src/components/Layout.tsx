import { Outlet } from 'react-router-dom';
import { Navigation } from './Navigation';

export function Layout() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-950 via-slate-900 to-slate-800 relative">
      <div className="absolute inset-0 bg-gradient-to-br from-slate-950/50 via-slate-900/30 to-slate-800/20" />
      <div className="absolute inset-0 bg-[radial-gradient(circle_at_30%_20%,rgba(251,191,36,0.1),transparent_60%)] opacity-60" />
      <div className="absolute inset-0 bg-[radial-gradient(circle_at_70%_80%,rgba(251,191,36,0.05),transparent_60%)] opacity-40" />
      
      <Navigation />
      <main className="pt-20 pb-8 relative z-10">
        <div className="container mx-auto px-4">
          <Outlet />
        </div>
      </main>
    </div>
  );
}