import { Link, useLocation } from 'react-router-dom';
import { cn } from '@/lib/utils';

export function Navigation() {
  const location = useLocation();

  const navItems = [
    { href: '/', label: '🏠 ホーム' },
    { href: '/movies', label: '🎦 映画一覧' },
    { href: '/genre', label: '🏷️ ジャンルから選ぶ' },
    { href: '/search', label: '🔍 検索' },
  ];

  return (
    <nav className="fixed top-0 left-0 right-0 z-50 bg-white/95 backdrop-blur-md border-b border-white/20 shadow-lg">
      <div className="container mx-auto px-4">
        <div className="flex justify-between items-center py-4">
          <Link
            to="/"
            className="text-2xl font-bold text-indigo-600 hover:text-indigo-700 transition-colors"
          >
            🎬 Go Movie Explorer
          </Link>
          
          <ul className="hidden md:flex gap-8">
            {navItems.map((item) => (
              <li key={item.href}>
                <Link
                  to={item.href}
                  className={cn(
                    'text-gray-700 hover:text-indigo-600 font-medium transition-colors relative',
                    location.pathname === item.href && 'text-indigo-600'
                  )}
                >
                  {item.label}
                  {location.pathname === item.href && (
                    <span className="absolute -bottom-1 left-0 right-0 h-0.5 bg-indigo-600 rounded-full" />
                  )}
                </Link>
              </li>
            ))}
          </ul>

          {/* Mobile menu button - 将来的に実装予定 */}
          <button className="md:hidden p-2 text-gray-600 hover:text-gray-800">
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M4 6h16M4 12h16M4 18h16" />
            </svg>
          </button>
        </div>
      </div>
    </nav>
  );
}