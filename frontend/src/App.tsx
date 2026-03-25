import React, { useState } from 'react';
import { CheckSquare, LogOut } from 'lucide-react';
import Login from './pages/Login';
import Signup from './pages/Signup';
import Dashboard from './pages/Dashboard';

export default function App() {
  const [token, setToken] = useState<string | null>(localStorage.getItem('token'));
  const [currentView, setCurrentView] = useState<'login' | 'signup' | 'dashboard'>(
    localStorage.getItem('token') ? 'dashboard' : 'login'
  );

  const login = (newToken: string) => {
    localStorage.setItem('token', newToken);
    setToken(newToken);
    setCurrentView('dashboard');
  };

  const logout = () => {
    localStorage.removeItem('token');
    setToken(null);
    setCurrentView('login');
  };

  return (
    <div className="min-h-screen bg-white text-black font-sans selection:bg-blue-600 selection:text-white">
      <nav className="border-b-4 border-black p-4 flex justify-between items-center">
        <div className="flex items-center gap-2 font-black text-2xl tracking-tight">
          <CheckSquare size={28} className="text-blue-600" />
          <span>Task<span className="text-blue-600">Manager</span></span>
        </div>
        
        {token && (
          <button 
            onClick={logout}
            className="flex items-center gap-2 font-bold text-sm border-2 border-black px-4 py-2 hover:bg-black hover:text-white transition-colors"
          >
            Logout <LogOut size={16} />
          </button>
        )}
      </nav>

      <main className="p-4 md:p-8 lg:p-12 max-w-6xl mx-auto">
        {currentView === 'login' && <Login onLogin={login} onNavigate={() => setCurrentView('signup')} />}
        {currentView === 'signup' && <Signup onNavigate={() => setCurrentView('login')} />}
        {currentView === 'dashboard' && <Dashboard />}
      </main>
    </div>
  );
}