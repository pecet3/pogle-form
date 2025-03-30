import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAppContext } from "../contexts/AppContext";
import { Loading } from "../components/Loading";
import { Aside } from "../components/Aside";

export const ProtectedPage = ({ children }: { children: React.ReactNode }) => {
  const { user, setUser } = useAppContext();
  const navigate = useNavigate();

  const fetchUser = async () => {
    const result = await fetch("/api/auth/ping");
    if (result.ok) {
      setUser(true);
    } else {
      navigate("/auth");
    }
  };

  useEffect(() => {
    if (!user) {
      fetchUser();
    }
  }, []);
  return (
    <>
      {user ? (
        <div className="flex">
          <Aside />
          <main>{children}</main>
        </div>
      ) : (
        <div className="my-40">
          <Loading />
        </div>
      )}
    </>
  );
};
