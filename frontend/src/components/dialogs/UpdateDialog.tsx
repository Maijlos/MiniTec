import { useState } from "react";
import { projectKeys } from "../../api/keys";
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import { getProject, updateProject } from "../../api/projects";

export function UpdateDialog({ id }: { id: number }) {
  const { data, isError, isLoading } = useQuery({
    queryKey: projectKeys.projectInfo(id),
    queryFn: () => getProject(id),
  });

  const loading = () => <div>Loading...</div>;

  const error = () => <div>Something went wrong!</div>;

  return (
    <>
      <dialog id={`update-dialog-${id}`} className="modal">
        <div className="modal-box">
          <div className="text-4xl ml-7">Update Project</div>
          {isLoading ? (
            loading()
          ) : isError ? (
            error()
          ) : (
            <Form
              initialName={data.data[0].name}
              initialCode={data.data[0].code}
              id={id}
            />
          )}
        </div>
        <form method="dialog" className="modal-backdrop">
          <button>close</button>
        </form>
      </dialog>
    </>
  );
}

function Form({
  id,
  initialName,
  initialCode,
}: {
  id: number;
  initialName: string;
  initialCode: string;
}) {
  const [name, setName] = useState(initialName);
  const [code, setCode] = useState(initialCode);
  const isDisabled = code === "" ? "disabled" : "";

  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: () => updateProject(id, name, code),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: projectKeys.allProjects });
      queryClient.invalidateQueries({ queryKey: projectKeys.projectInfo(id) });
    },
  });

  const handleSubmit = () => {
    mutation.mutate();
    const modal = document.getElementById(
      `update-dialog-${id}`,
    ) as HTMLDialogElement | null;
    modal?.close();
  };

  return (
    <>
      <div className="m-9">
        <label className="floating-label">
          <input
            type="text"
            placeholder="Name"
            className="input w-full input-xl"
            value={name}
            onChange={(e) => setName(e.target.value)}
          />
          <span>Name</span>
        </label>
        <label className="floating-label mt-6">
          <input
            type="text"
            placeholder="Code"
            className="input w-full input-xl required:"
            value={code}
            onChange={(e) => setCode(e.target.value)}
          />
          <span>Code</span>
        </label>
      </div>
      <div className="flex justify-center">
        <button
          type="submit"
          className={`btn ${isDisabled}  text-3xl`}
          disabled={code === ""}
          onClick={handleSubmit}
        >
          Update
        </button>
      </div>
    </>
  );
}
