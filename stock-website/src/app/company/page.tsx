import dynamic from "next/dynamic";
import { columns, Payment } from "./Sections/FinancialStatements/columns";
import { DataTable } from "./Sections/FinancialStatements/data-table";
import ValuationMetrics from "./ValuationMetrics";

const SimpleChart = dynamic(() => import("./SimpleChart"));

async function getData(): Promise<Payment[]> {
  // Fetch data from your API here.
  return [
    {
      id: "728ed52f",
      amount: 100,
      status: "pending",
      email: "m@example.com",
    },
    // ...
  ];
}

export default async function CompanyPage() {
  const data = await getData();

  return (
    <>
      <div>
        <h1>Company Page</h1>
      </div>
      <div className="flex flex-row">
        <div className="w-1/2">
          <ValuationMetrics />
        </div>
        <div className="w-1/2">
          <SimpleChart />
        </div>
      </div>

      <div>
        {/* Table */}
        <DataTable columns={columns} data={data} />
      </div>
    </>
  );
}
