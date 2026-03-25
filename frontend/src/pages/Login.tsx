import React, { useState } from 'react';
import { ArrowRight, AlertTriangle } from 'lucide-react';
import api from '../api/axios';

interface LoginProps {
  onLogin: (token: string) => void;
  onNavigate: () => void;
}

export default function Login({ onLogin, onNavigate }: LoginProps) {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError('');
    setLoading(true);
    try {
      const res = await api.post('/auth/login', { email, password });
      if (res.data.success) {
        onLogin(res.data.data.token);
      } else {
        setError(res.data.error || 'Login failed');
      }
    } catch (err: any) {
      setError(err.response?.data?.error || 'Invalid email or password');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="max-w-md mx-auto mt-12">
      <div className="border-4 border-black p-8 shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] bg-white">
        <h1 className="text-4xl font-black mb-2 tracking-tight">Welcome Back</h1>
        <p className="text-gray-600 font-medium mb-8">Please sign in to your account.</p>

        {error && (
          <div className="bg-red-500 text-white p-3 mb-6 font-bold flex items-center gap-2 border-2 border-black">
            <AlertTriangle size={20} /> {error}
          </div>
        )}

        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label className="block font-bold mb-2 text-sm">Email Address</label>
            <input 
              type="email" 
              required
              className="w-full border-2 border-black p-3 focus:outline-none focus:border-blue-600 focus:ring-0 transition-colors"
              value={email} onChange={e => setEmail(e.target.value)} 
            />
          </div>
          <div>
            <label className="block font-bold mb-2 text-sm">Password</label>
            <input 
              type="password" 
              required
              className="w-full border-2 border-black p-3 focus:outline-none focus:border-blue-600 focus:ring-0 transition-colors"
              value={password} onChange={e => setPassword(e.target.value)} 
            />
          </div>
          <button 
            disabled={loading}
            className="w-full bg-black text-white font-bold text-lg p-4 flex justify-between items-center hover:bg-blue-600 transition-colors disabled:opacity-50"
          >
            {loading ? 'Signing in...' : 'Sign In'} <ArrowRight />
          </button>
        </form>

        <div className="mt-8 pt-6 border-t-2 border-black text-center">
          <span className="font-medium text-gray-600">Don't have an account? </span>
          <button onClick={onNavigate} className="font-bold text-blue-600 hover:underline">
            Sign up here
          </button>
        </div>
      </div>
    </div>
  );
}