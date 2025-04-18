import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAppContext } from "../contexts/AppContext";
import { Loading } from "../components/Loading";
import { Aside } from "../components/Aside";

export const ProtectedPage = ({ children }: { children: React.ReactNode }) => {
  const { user, setUser, courses, setCourses, persons, setPersons } =
    useAppContext();
  const navigate = useNavigate();

  const fetchUser = async () => {
    const result = await fetch("/api/auth/ping");
    if (result.ok) {
      setUser(true);
    } else {
      navigate("/auth");
    }
  };

  const fetchCourses = async () => {
    const result = await fetch("/api/courses");
    if (result.ok) {
      const data = await result.json();
      setCourses(data);
    } else {
      navigate("/auth");
    }
  };

  const fetchPersons = async () => {
    const result = await fetch("/api/persons");
    if (result.ok) {
      const data = await result.json();
      setPersons(data);
      console.log(data);
    } else {
      navigate("/auth");
    }
  };

  useEffect(() => {
    if (!user) {
      fetchUser();
    }
    if (courses.length === 0) {
      fetchCourses();
    }
    if (persons.length === 0) {
      fetchPersons();
    }
  }, []);
  return (
    <>
      {user ? (
        <div className="flex">
          <Aside />
          <main className=" mx-auto flex flex-col justify-start">
            {children}
          </main>
        </div>
      ) : (
        <div className="my-40">
          <Loading />
        </div>
      )}
    </>
  );
};
