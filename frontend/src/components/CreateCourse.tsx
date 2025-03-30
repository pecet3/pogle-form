import { useState } from "react";

export const CreateCourse = () => {
  const [name, setName] = useState("");
  const [maxPersons, setMaxPersons] = useState("");
  const [error, setError] = useState("");
  const [successMessage, setSuccessMessage] = useState("");

  const handleSubmit = async (event: any) => {
    event.preventDefault();
    setError("");
    setSuccessMessage("");

    if (!name.trim()) {
      setError("Nazwa kursu jest wymagana.");
      return;
    }

    const maxPersonsNumber = parseInt(maxPersons, 10);
    if (isNaN(maxPersonsNumber) || maxPersonsNumber <= 0) {
      setError("Maksymalna liczba osób musi być liczbą większą od 0.");
      return;
    }

    try {
      const response = await fetch("/api/courses", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ name, max_persons: maxPersonsNumber }),
      });

      if (response.ok) {
        setSuccessMessage("Kurs został utworzony pomyślnie!");
        setName("");
        setMaxPersons("");
      } else {
        const errorData = await response.json();
        setError(errorData.message || "Wystąpił błąd podczas tworzenia kursu.");
      }
    } catch (error) {
      setError("Wystąpił błąd podczas komunikacji z serwerem.");
      console.error("Błąd:", error);
    }
  };

  return (
    <div className="">
      <h2>Stwórz nowy kurs</h2>
      {error && <p className="error">{error}</p>}
      {successMessage && <p className="success">{successMessage}</p>}
      <form
        onSubmit={handleSubmit}
        className="cartoon-container max-w-md mx-auto mt-4"
      >
        <div className="mb-4">
          <label
            htmlFor="name"
            className="block text-gray-700 text-sm font-bold mb-2"
          >
            Nazwa kursu:
          </label>
          <input
            type="text"
            id="name"
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
        </div>
        <div className="mb-4">
          <label
            htmlFor="maxPersons"
            className="block text-gray-700 text-sm font-bold mb-2"
          >
            Maksymalna liczba osób:
          </label>
          <input
            type="number"
            id="maxPersons"
            className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
            value={maxPersons}
            onChange={(e) => setMaxPersons(e.target.value)}
          />
        </div>
        <div className="flex items-center justify-between">
          <button
            className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
            type="submit"
          >
            Stwórz kurs
          </button>
        </div>
      </form>
    </div>
  );
};
