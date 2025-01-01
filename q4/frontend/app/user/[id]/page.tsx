"use client";

import React from "react";
import UserForm from "@/components/UserForm";
import { useParams } from "next/navigation";

const EditUser: React.FC = () => {
  const params = useParams();
  const id = Array.isArray(params?.id) ? params.id[0] : params?.id; // Ensure id is a string

  return <UserForm id={id} />;
};

export default EditUser;
