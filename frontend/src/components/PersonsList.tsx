import { useAppContext } from "../contexts/AppContext";

export const PersonsList = () => {
  const { persons } = useAppContext();
  return (
    <div>
      <h2>Lista zarejestrowanych osób</h2>
      {persons.length === 0 ? (
        <p>Brak zarejestrowanych osób.</p>
      ) : (
        <ul className="list-disc pl-5">
          {persons.map((person) => (
            <li key={person.id} className="mb-2">
              <strong>{person.full_name}</strong> ({person.email})
              {person.chosen_course && (
                <p className="text-sm text-gray-600">
                  Wybrany kurs: {person.chosen_course.name}
                </p>
              )}
              {!person.chosen_course && (
                <p className="text-sm text-gray-600">Wybrany kurs: Brak</p>
              )}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

interface PersonsByCourse {
  course: string;
  persons: Person[];
}

export const PersonsListSortedByCourse = () => {
  const { persons } = useAppContext();

  // Grupowanie osób według nazwy kursu
  const groupedPersons = persons.reduce(
    (acc: Record<string, Person[]>, person) => {
      const courseName = person.chosen_course
        ? person.chosen_course.name
        : "Brak kursu";
      if (!acc[courseName]) {
        acc[courseName] = [];
      }
      acc[courseName].push(person);
      return acc;
    },
    {}
  );

  // Konwersja zgrupowanych osób do tablicy obiektów z nazwą kursu i listą osób
  const sortedPersonsByCourse: PersonsByCourse[] = Object.entries(
    groupedPersons
  )
    .sort(([courseNameA], [courseNameB]) =>
      courseNameA.localeCompare(courseNameB)
    ) // Sortowanie alfabetyczne po nazwie kursu
    .map(([courseName, personsList]) => ({
      course: courseName,
      persons: personsList,
    }));

  return (
    <div>
      <h2>Lista osób posortowana według kursu</h2>
      {sortedPersonsByCourse.length === 0 ? (
        <p>Brak zarejestrowanych osób.</p>
      ) : (
        <div>
          {sortedPersonsByCourse.map((group) => (
            <div key={group.course} className="mb-4">
              <h3>{group.course}</h3>
              {group.persons.length === 0 ? (
                <p className="text-sm text-gray-500">
                  Brak osób zapisanych na ten kurs.
                </p>
              ) : (
                <ul className="list-disc pl-5">
                  {group.persons.map((person) => (
                    <li key={person.id} className="mb-1">
                      {person.full_name} ({person.email})
                    </li>
                  ))}
                </ul>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
};
