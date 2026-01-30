import type { SuccessfulData } from "../api/projects.tsx";

type DropdownMenuProps = {
  projects: SuccessfulData;
  setChosedProject: (id: number) => void;
};

export function DropdownMenu({
  projects,
  setChosedProject,
}: DropdownMenuProps) {
  return (
    <div className="dropdown">
      <div tabIndex={0} role="button" className="btn m-5 p-9 text-4xl bg-amber-300">
        Choose Project
      </div>
      <ul
        tabIndex={0}
        className="dropdown-content menu bg-base-100 text-2xl rounded-box z-1 w-52 p-2 shadow-sm"
      >
        {Array.isArray(projects?.data) && projects.data.map((project) => (
          <li key={project.id}>
            <a onClick={() => setChosedProject(project.id)}>{project.name}</a>
          </li>
        ))}
      </ul>
    </div>
  );
}
