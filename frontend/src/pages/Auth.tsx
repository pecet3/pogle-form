import React, { useState } from "react";

interface Login {
  name: string;
  password: string;
}

export const Auth: React.FC = () => {
  const [name, setName] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    setError(null);
    setLoading(true);

    const loginData: Login = {
      name: name,
      password: password,
    };

    try {
      const response = await fetch("/api/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(loginData),
      });

      if (response.ok) {
        // Obsłuż pomyślne logowanie, np. przekierowanie użytkownika
        console.log("Zalogowano pomyślnie");
        // Możesz tutaj dodać logikę przekierowania, np. używając `window.location.href` lub biblioteki do routingu
      } else {
        // Obsłuż błąd logowania
        const errorData = await response.json();
        setError(errorData.message || "Błąd logowania");
      }
    } catch (err) {
      console.error("Wystąpił błąd podczas wysyłania żądania:", err);
      setError("Wystąpił błąd podczas logowania. Spróbuj ponownie później.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4 w-96 m-auto flex flex-col items-center"
    >
      {error && <p className="text-red-500 text-xs italic mb-4">{error}</p>}
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="name"
        >
          Nazwa użytkownika:
        </label>
        <input
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="name"
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
      </div>
      <div className="mb-6">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="password"
        >
          Hasło:
        </label>
        <input
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="password"
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
      </div>
      <div className="flex items-center justify-between m-auto">
        <button
          className={`bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline ${
            loading ? "opacity-50 cursor-not-allowed" : ""
          }`}
          type="submit"
          disabled={loading}
        >
          {loading ? "Logowanie..." : "Zaloguj się"}
        </button>
      </div>
    </form>
  );
};
