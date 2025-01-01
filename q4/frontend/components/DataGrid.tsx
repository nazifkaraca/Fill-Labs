"use client";

import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import axios from "axios";

interface User {
    id: number;
    name: string;
    email: string;
    age: number;
}

const DataGrid: React.FC = () => {
    const [users, setUsers] = useState<User[]>([]);
    const router = useRouter();

    useEffect(() => {
        axios
            .get("http://localhost:8080/users")
            .then((response) => setUsers(response.data))
            .catch((error) => console.error("Error fetching users:", error));
    }, []);

    const handleDelete = (id: number) => {
        if (confirm("Are you sure you want to delete this user?")) {
            axios
                .delete(`http://localhost:8080/users/${id}`)
                .then(() => setUsers(users.filter((user) => user.id !== id)))
                .catch((error) => console.error("Error deleting user:", error));
        }
    };

    return (
        <div>
            <button
                onClick={() => router.push("/user/new")}
                style={{
                    marginBottom: "16px",
                    padding: "8px 16px",
                    backgroundColor: "#007bff",
                    color: "#fff",
                    border: "none",
                    borderRadius: "4px",
                    cursor: "pointer",
                }}
            >
                New User
            </button>
            <table
                style={{
                    width: "100%",
                    borderCollapse: "collapse",
                    marginTop: "16px",
                    fontSize: "16px",
                }}
            >
                <thead>
                    <tr style={{ backgroundColor: "#f9f9f9", textAlign: "left" }}>
                        <th style={{ padding: "8px", border: "1px solid #ddd" }}>ID</th>
                        <th style={{ padding: "8px", border: "1px solid #ddd" }}>Name</th>
                        <th style={{ padding: "8px", border: "1px solid #ddd" }}>Email</th>
                        <th style={{ padding: "8px", border: "1px solid #ddd" }}>Age</th>
                        <th style={{ padding: "8px", border: "1px solid #ddd" }}>Actions</th>
                    </tr>
                </thead>
                <tbody>
                    {users.map((user, index) => (
                        <tr key={user.id} style={{ borderBottom: "1px solid #ddd" }}>
                            <td style={{ padding: "8px" }}>{index + 1}</td>
                            <td style={{ padding: "8px" }}>{user.name}</td>
                            <td style={{ padding: "8px" }}>{user.email}</td>
                            <td style={{ padding: "8px" }}>{user.age}</td>
                            <td style={{ padding: "8px" }}>
                                <button
                                    onClick={() => router.push(`/user/${user.id}`)}
                                    style={{
                                        marginRight: "8px",
                                        padding: "4px 8px",
                                        backgroundColor: "#ffc107",
                                        border: "none",
                                        borderRadius: "4px",
                                        cursor: "pointer",
                                        color: "#fff",
                                    }}
                                >
                                    Edit
                                </button>
                                <button
                                    onClick={() => handleDelete(user.id)}
                                    style={{
                                        padding: "4px 8px",
                                        backgroundColor: "#dc3545",
                                        border: "none",
                                        borderRadius: "4px",
                                        cursor: "pointer",
                                        color: "#fff",
                                    }}
                                >
                                    Delete
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
};

export default DataGrid;
