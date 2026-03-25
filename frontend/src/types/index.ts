export interface Task {
  id: number;
  title: string;
  description: string;
  status: string;
  user_id: number;
  created_at: string;
  updated_at: string;
  deleted_at: string | null;
}