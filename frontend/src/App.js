import React, { useEffect, useState } from 'react';

// Simple React app that reads REACT_APP_API_URL from environment.
// During development, .env.development sets REACT_APP_API_URL to http://localhost:8080
// In production it should be set to https://api.monsite.com
const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:10000';

export default function App() {
  const [employees, setEmployees] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    // Fetch employees from backend API
    fetch(`${API_URL}/employees`)
      .then(res => {
        if (!res.ok) throw new Error('Network response was not ok');
        return res.json();
      })
      .then(body => setEmployees(body.data || []))
      .catch(err => setError(err.message));
  }, []);

  return (
    <div style={{ padding: 20, fontFamily: 'Arial, sans-serif' }}>
      <h1>Gestion de paie — Exemple</h1>
      <p>API utilisée: <code>{API_URL}</code></p>
      {error && <div style={{ color: 'red' }}>Erreur: {error}</div>}
      <ul>
        {employees.map(emp => (
          <li key={emp.id}>
            {emp.name} — {emp.role} — Salaire: {emp.salary}
          </li>
        ))}
      </ul>
    </div>
  );
}
