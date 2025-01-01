"use client";

import React, { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import axios from "axios";

interface UserFormProps {
    id?: string;
}

const UserForm: React.FC<UserFormProps> = ({ id }) => {
    const [user, setUser] = useState({ name: "", email: "", age: "" });
    const router = useRouter();
    const [error, setError] = useState<string | null>(null);

    useEffect(() => {
        if (id) {
            axios.get(`http://localhost:8080/users/${id}`)
                .then((response) => setUser(response.data))
                .catch((error) => console.error("Error fetching user:", error));
        }
    }, [id]);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        if (!user.name || !user.email || !user.age) {
            setError("All fields are required.");
            return;
        }
        const request = id
            ? axios.put(`http://localhost:8080/users/${id}`, {
                name: user.name,
                email: user.email,
                age: parseInt(user.age),
            })
            : axios.post("http://localhost:8080/users", {
                name: user.name,
                email: user.email,
                age: parseInt(user.age),
            });

        request
            .then(() => router.push("/"))
            .catch((error) => console.error("Error saving user:", error));
    };

    return (
        <form onSubmit={handleSubmit} style={{ maxWidth: "400px", margin: "0 auto" }}>
            <h2>{id ? "Edit User" : "New User"}</h2>
            {error && <p style={{ color: "red" }}>{error}</p>}
            <div style={{ marginBottom: "16px" }}>
                <label>Name:</label>
                <input
                    type="text"
                    value={user.name}
                    onChange={(e) => setUser({ ...user, name: e.target.value })}
                    style={{
                        width: "100%",
                        padding: "8px",
                        border: "1px solid #ddd",
                        borderRadius: "4px",
                        marginTop: "4px",
                    }}
                    required
                />
            </div>
            <div style={{ marginBottom: "16px" }}>
                <label>Email:</label>
                <input
                    type="email"
                    value={user.email}
                    onChange={(e) => setUser({ ...user, email: e.target.value })}
                    style={{
                        width: "100%",
                        padding: "8px",
                        border: "1px solid #ddd",
                        borderRadius: "4px",
                        marginTop: "4px",
                    }}
                    required
                />
            </div>
            <div style={{ marginBottom: "16px" }}>
                <label>Age:</label>
                <input
                    type="number"
                    value={user.age}
                    onChange={(e) => setUser({ ...user, age: e.target.value })}
                    style={{
                        width: "100%",
                        padding: "8px",
                        border: "1px solid #ddd",
                        borderRadius: "4px",
                        marginTop: "4px",
                    }}
                    required
                />
            </div>
            <button
                type="submit"
                style={{
                    padding: "8px 16px",
                    backgroundColor: "#007bff",
                    color: "#fff",
                    border: "none",
                    borderRadius: "4px",
                    cursor: "pointer",
                }}
            >
                {id ? "Save" : "Create"}
            </button>
            <button
                type="button"
                onClick={() => router.push("/")}
                style={{
                    padding: "8px 16px",
                    backgroundColor: "#6c757d",
                    color: "#fff",
                    border: "none",
                    borderRadius: "4px",
                    cursor: "pointer",
                    marginLeft: "8px",
                }}
            >
                Back
            </button>
        </form>
    );
};

export default UserForm;
