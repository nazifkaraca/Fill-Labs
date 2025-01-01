import React from "react";
import DataGrid from "@/components/DataGrid";

const Home: React.FC = () => {
  return (
    <div style={{ padding: "20px" }}>
      <h1>User Management</h1>
      <DataGrid />
    </div>
  );
};

export default Home;
