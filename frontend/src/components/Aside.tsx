import { Link, useLocation } from "react-router-dom";

const pages = [
  { link: "/courses", title: "Warsztaty" },
  { link: "/persons", title: "Osoby" },
];

export const Aside = () => {
  const location = useLocation();

  return (
    <aside
      className="flex bg-white  flex-col justify-end items-center
     w-72 border-r h-screen"
    >
      <nav className="flex flex-col w-full">
        {pages.map((page) => (
          <Link
            to={page.link}
            key={page.link}
            className={`block w-full p-4 text-left ${
              location.pathname === page.link ? "bg-gray-200" : ""
            }`}
          >
            {page.title}
          </Link>
        ))}
      </nav>
      <div className="border-t w-full p-2">
        <div className="flex items-center justify-center gap-1 p-2">
          <Link to={"/"} className="font-mono text-2xl">
            Pogle-Form ADMIN
          </Link>
        </div>
      </div>
    </aside>
  );
};
