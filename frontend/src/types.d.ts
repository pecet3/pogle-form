interface Course {
  id: number;
  name: string;
  max_persons: number;
  registered_persons: number;
}

interface Person {
  id: number;
  full_name: string;
  email: string;
  chosen_course: Course | null;
}