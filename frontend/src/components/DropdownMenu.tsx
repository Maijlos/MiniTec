import { useMutation, useQueryClient } from "@tanstack/react-query";
import type { SuccessfulData, ErrorData } from "../api/projects.tsx";
import { deleteProject } from "../api/projects.tsx";
import { projectKeys } from "../api/keys.tsx";

type DropdownMenuProps = {
  projects: SuccessfulData | ErrorData | undefined;
  setChosedProject: (id: number) => void;
};

export function DropdownMenu({
  projects,
  setChosedProject,
}: DropdownMenuProps) {
  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: (id: number) => deleteProject(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: projectKeys.allProjects });
    },
  });

  return (
    <div className="dropdown">
      <div
        tabIndex={0}
        role="button"
        className="btn m-5 p-9 text-4xl bg-amber-300"
      >
        Choose Project
      </div>
      <ul
        tabIndex={-1}
        className="dropdown-content menu bg-base-100 text-2xl rounded-box z-1 w-70 p-2 ml-5 shadow-sm"
      >
        {projects && 'data' in projects && Array.isArray(projects.data) && projects.data.length > 0 ? 
        (projects.data.map((project) => (
            <li key={project.id}>
              <div className="flex justify-between">
                <a className="" onClick={() => setChosedProject(project.id)}>
                  {project.name}
                </a>
                <a
                  className="text-sm flex-none hover:bg-red-300 rounded-sm text-red-500"
                  onClick={() => mutation.mutate(project.id)}
                >
                  Delete
                </a>
              </div>
            </li>
          ))) : (<li>No projects found!</li>)}
      </ul>
    </div>
  );
}
