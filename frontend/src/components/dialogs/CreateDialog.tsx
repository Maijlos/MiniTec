import { useState } from "react";
import { projectKeys } from "../../api/keys";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { createProject } from "../../api/projects";

export function CreateDialog() {
  const [name, setName] = useState("");
  const [code, setCode] = useState("");
  const isDisabled = code === "" ? "disabled" : "";

  const queryClient = useQueryClient();

  const mutation = useMutation({
    mutationFn: () => createProject(name, code),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: projectKeys.allProjects });
    },
  });

  const handleSubmit = () => {
    mutation.mutate();
    const modal = document.getElementById(
      "create-update-dialog",
    ) as HTMLDialogElement | null;
    modal?.close();
    setName("");
    setCode("");
  };

  return (
    <>
      <dialog id="create-dialog" className="modal">
        <div className="modal-box">
          <div className="text-4xl ml-7">Create Project</div>
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
              Create
            </button>
          </div>
        </div>
        <form method="dialog" className="modal-backdrop">
          <button>close</button>
        </form>
      </dialog>
    </>
  );
}
