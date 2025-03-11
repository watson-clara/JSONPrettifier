import React, { useState } from 'react';
import axios from 'axios';
import './JsonPrettifier.css';

const JsonPrettifier: React.FC = () => {
  const [inputJson, setInputJson] = useState('');
  const [outputJson, setOutputJson] = useState('');
  const [error, setError] = useState('');

  const handleInputChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setInputJson(e.target.value);
  };

  const formatJSON = async () => {
    try {
      const query = `
        query {
          formatJSON(input: ${JSON.stringify(inputJson)})
        }
      `;
      
      const response = await axios.post('http://localhost:8080/graphql', {
        query
      });
      
      if (response.data.errors) {
        throw new Error(response.data.errors[0].message);
      }
      
      setOutputJson(response.data.data.formatJSON);
      setError('');
    } catch (err) {
      setError('Invalid JSON input');
      setOutputJson('');
    }
  };

  return (
    <div className="json-prettifier">
      <h1>JSON Prettifier</h1>
      
      <div className="input-container">
        <h2>Input JSON</h2>
        <textarea
          value={inputJson}
          onChange={handleInputChange}
          placeholder="Paste your JSON here..."
          rows={10}
        />
      </div>
      
      <button onClick={formatJSON}>Format JSON</button>
      
      {error && <div className="error">{error}</div>}
      
      <div className="output-container">
        <h2>Formatted JSON</h2>
        <pre>{outputJson}</pre>
      </div>
    </div>
  );
};

export default JsonPrettifier; 