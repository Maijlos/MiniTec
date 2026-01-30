import { useQuery } from "@tanstack/react-query";
import { projectKeys } from "../api/keys";
import { getProjectHealth } from "../api/projects";

export function Table({ id }: { id: number }) {
  const {
    data: projectHealth,
    isError,
    isLoading,
  } = useQuery({
    queryKey: projectKeys.projectInfo(id),
    queryFn: () => getProjectHealth(id),
  });

  if (isError) {
    return <div>Something went wrong!</div>;
  }

  if (isLoading) {
    return <div>Loading...</div>;
  }

  console.log(projectHealth);

  if (projectHealth && "data" in projectHealth && projectHealth.data) {
    let keyCount = 0;
    const rows = [];
    for (
      let i = 0;
      i < projectHealth.data[Object.keys(projectHealth.data)[0]].length;
      i++
    ) {
      const row = writeRow(projectHealth, keyCount);
      keyCount++;
      rows.push(row);
    }

    console.log(rows);

    return (
      <div className="overflow-x-auto">
        <table className="table table-lg w-max table-zebra">
          <thead>
            <tr>
              {Object.keys(projectHealth.data).map((stationName) => {
                return (
                  <>
                    <th></th>
                    <th key={stationName}>{stationName}</th>
                    <th className="border-r"></th>
                  </>
                );
              })}
            </tr>
            <tr>
              {Object.keys(projectHealth.data).flatMap((stationName) => [
                <th key={`${stationName}-start`}>start date</th>,
                <th key={`${stationName}-end`}>end date</th>,
                <th key={`${stationName}-status`} className="border-r">
                  final status
                </th>,
              ])}
            </tr>
          </thead>
          <tbody>
            {rows.map((row) => {
              return (
                <tr>
                  {row.map((cell, cellIndex) => (
                  <td key={cellIndex} className={cellIndex % 3 === 2 ? "border-r" : ""}>
                    {cell}
                  </td>
                ))}
                </tr>
              );
            })}
          </tbody>
        </table>
      </div>
    );
  }
}

type StationData = {
  start_date: string;
  end_date: string;
  final_status: string;
};

type ProjectHealthData = {
  data: {
    [stationName: string]: StationData[];
  };
};

function writeRow(projectHealth: ProjectHealthData, keyCount: number) {
  const row = Object.values(projectHealth.data).flatMap((arrays) => [
    arrays[keyCount].start_date,
    arrays[keyCount].end_date,
    arrays[keyCount].final_status,
  ]);

  return row;
}
