export interface Manager {
    id: string;
    name: string;
    grade: string;
    password: string;
    ttl: number | null;
    create_at: string;  // or Date if you want to parse it later
    update_at: string;
  }