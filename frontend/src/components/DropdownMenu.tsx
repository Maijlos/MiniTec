import type { SuccessfulData } from "../api/projects.tsx";

type DropdownMenuProps = {
  projects: SuccessfulData | undefined;
  setChosedProject: (id: number) => void;
};

export function DropdownMenu({
  projects,
  setChosedProject,
}: DropdownMenuProps) {
  return (
    <div className="dropdown">
      <div tabIndex={0} role="button" className="btn m-1">
        Choose Project
      </div>
      <ul
        tabIndex={0}
        className="dropdown-content menu bg-base-100 rounded-box z-1 w-52 p-2 shadow-sm"
      >
        {projects?.data?.map((project) => (
          <li key={project.id}>
            <a onClick={() => setChosedProject(project.id)}>{project.name}</a>
          </li>
        ))}
      </ul>
    </div>
  );
}
