import React, { createContext, useContext, useState, ReactNode } from "react";

type AppContextType = {
  user: boolean;
  setUser: React.Dispatch<React.SetStateAction<boolean>>;
  courses: Course[] | null;
  setCourses: React.Dispatch<React.SetStateAction<Course[] | null>>;
};

const AppContext = createContext<AppContextType | undefined>(undefined);

export const AppContextProvider: React.FC<{ children: ReactNode }> = ({
  children,
}) => {
  const [user, setUser] = useState<boolean>(false);
  const [courses, setCourses] = useState<Course[] | null>(null);
  return (
    <AppContext.Provider
      value={{
        user,
        setUser,
        courses,
        setCourses,
      }}
    >
      {children}
    </AppContext.Provider>
  );
};

export const useAppContext = () => {
  const context = useContext(AppContext);
  if (context === undefined) {
    throw new Error("useAppContext must be used within a AppContext");
  }
  return context;
};
