import { useAppContext } from "../contexts/AppContext";

export const CoursesList = () => {
  const { courses } = useAppContext();

  return (
    <div>
      <h2>Lista dostępnych kursów</h2>
      {courses?.length === 0 ? (
        <p>Brak dostępnych kursów.</p>
      ) : (
        <ul className="list-disc pl-5">
          {courses?.map((course) => (
            <li key={course.id} className="mb-2">
              <strong>{course.name}</strong> ({course.registered_persons}/
              {course.max_persons} osób)
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};
