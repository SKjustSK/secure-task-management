import React, { useState, useEffect } from 'react';
import { Plus, Trash2, Calendar, Edit2, X, Save } from 'lucide-react';
import api from '../api/axios';
import type { Task } from '../types';

export default function Dashboard() {
  const [tasks, setTasks] = useState<Task[]>([]);
  const [loading, setLoading] = useState(true);

  const [title, setTitle] = useState('');
  const [desc, setDesc] = useState('');
  const [status, setStatus] = useState('pending');
  const [editingId, setEditingId] = useState<number | null>(null);

  const fetchTasks = async () => {
    try {
      const res = await api.get('/tasks');
      setTasks(res.data.data || []);
    } catch (err) {
      console.error("Fetch error", err);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => { fetchTasks(); }, []);

  const addTask = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!title.trim()) return;
    try {
      await api.post('/tasks', { title, description: desc, status: 'pending' });
      resetForm();
      fetchTasks();
    } catch (err) { console.error(err); }
  };

  const handleEditClick = async (id: number) => {
    try {
      const res = await api.get(`/tasks/${id}`);
      const task = res.data.data;
      
      setEditingId(task.id);
      setTitle(task.title);
      setDesc(task.description);
      setStatus(task.status);
    } catch (err) {
      console.error("Failed to fetch task details", err);
    }
  };

  const updateTask = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!editingId || !title.trim()) return;
    try {
      await api.put(`/tasks/${editingId}`, { title, description: desc, status });
      resetForm();
      fetchTasks();
    } catch (err) { console.error(err); }
  };

  const deleteTask = async (id: number) => {
    try {
      await api.delete(`/tasks/${id}`);
      if (editingId === id) resetForm();
      fetchTasks();
    } catch (err) { console.error(err); }
  };

  const resetForm = () => {
    setEditingId(null);
    setTitle('');
    setDesc('');
    setStatus('pending');
  };

  const formatDate = (dateString: string) => {
    if (!dateString) return '';
    return new Date(dateString).toLocaleDateString('en-US', { 
      month: 'short', day: 'numeric', year: 'numeric' 
    });
  };

  return (
    <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
      {/* Sidebar Form */}
      <div className="lg:col-span-1">
        <div className={`border-4 border-black p-6 shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] sticky top-8 transition-colors ${editingId ? 'bg-yellow-50' : 'bg-white'}`}>
          
          <div className="flex justify-between items-center mb-6 border-b-2 border-black pb-2">
            <h2 className="text-2xl font-black">
              {editingId ? 'Edit Task' : 'Create Task'}
            </h2>
            {editingId && (
              <button onClick={resetForm} className="text-gray-500 hover:text-red-600 transition-colors">
                <X size={24} />
              </button>
            )}
          </div>

          <form onSubmit={editingId ? updateTask : addTask} className="space-y-4">
            <div>
              <label className="block font-bold mb-2 text-sm">Task Title</label>
              <input 
                type="text" 
                required
                className="w-full border-2 border-black p-3 focus:outline-none focus:border-blue-600 transition-colors bg-white"
                value={title} onChange={e => setTitle(e.target.value)} 
              />
            </div>
            
            <div>
              <label className="block font-bold mb-2 text-sm">Description (Optional)</label>
              <textarea 
                className="w-full border-2 border-black p-3 focus:outline-none focus:border-blue-600 transition-colors h-24 resize-none bg-white"
                value={desc} onChange={e => setDesc(e.target.value)} 
              />
            </div>

            {editingId && (
              <div>
                <label className="block font-bold mb-2 text-sm">Status</label>
                <select 
                  className="w-full border-2 border-black p-3 focus:outline-none focus:border-blue-600 transition-colors bg-white font-bold cursor-pointer"
                  value={status} onChange={e => setStatus(e.target.value)}
                >
                  <option value="pending">Pending</option>
                  <option value="in_progress">In Progress</option>
                  <option value="completed">Completed</option>
                </select>
              </div>
            )}

            <button 
              type="submit"
              className={`w-full text-white font-bold text-lg p-3 flex justify-center items-center gap-2 hover:bg-black transition-colors border-2 border-transparent hover:border-black ${editingId ? 'bg-yellow-600' : 'bg-blue-600'}`}
            >
              {editingId ? <><Save size={20} /> Save Changes</> : <><Plus size={20} /> Add Task</>}
            </button>
          </form>
        </div>
      </div>

      {/* Main Task List */}
      <div className="lg:col-span-2 space-y-4">
        <h2 className="text-3xl font-black mb-6 flex items-center gap-4">
          My Tasks
          <span className="bg-black text-white text-sm py-1 px-3 rounded-full">{tasks.length}</span>
        </h2>

        {loading ? (
          <div className="p-8 border-4 border-black border-dashed font-bold text-center text-gray-500 animate-pulse">
            Loading tasks...
          </div>
        ) : tasks.length === 0 ? (
          <div className="p-12 border-4 border-black border-dashed font-bold text-center text-gray-500">
            No tasks yet. Create one to get started.
          </div>
        ) : (
          tasks.map(task => (
            <div 
              key={task.id} 
              className={`border-4 border-black p-5 bg-white shadow-[4px_4px_0px_0px_rgba(0,0,0,1)] flex justify-between items-start group hover:-translate-y-1 hover:shadow-[8px_8px_0px_0px_rgba(0,0,0,1)] transition-all ${editingId === task.id ? 'ring-4 ring-yellow-400 border-yellow-500' : ''}`}
            >
              <div className="w-full pr-4">
                <h3 className={`text-xl font-bold ${task.status === 'completed' ? 'line-through text-gray-400' : ''}`}>
                  {task.title}
                </h3>
                
                <div className="flex items-center gap-1 text-gray-500 font-bold text-xs mt-1 mb-3">
                  <Calendar size={14} /> 
                  <span>Created: {formatDate(task.created_at)}</span>
                </div>

                {task.description && <p className={`font-medium ${task.status === 'completed' ? 'text-gray-400' : 'text-gray-700'}`}>{task.description}</p>}
                
                <div className={`mt-4 inline-block border-2 border-black px-2 py-1 text-xs font-bold uppercase tracking-wider ${task.status === 'completed' ? 'bg-green-100 text-green-900' : task.status === 'in_progress' ? 'bg-yellow-100 text-yellow-900' : 'bg-blue-100 text-blue-900'}`}>
                  {task.status === 'in_progress' ? 'IN PROGRESS' : task.status.toUpperCase()}
                </div>
              </div>

              <div className="flex flex-col gap-2 shrink-0">
                <button 
                  onClick={() => handleEditClick(task.id)}
                  className="border-2 border-black p-2 text-black hover:bg-yellow-400 transition-colors"
                  title="Edit Task"
                >
                  <Edit2 size={20} />
                </button>
                <button 
                  onClick={() => deleteTask(task.id)}
                  className="border-2 border-black p-2 text-black hover:bg-red-500 hover:text-white transition-colors"
                  title="Delete Task"
                >
                  <Trash2 size={20} />
                </button>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
}