import React, { useState } from 'react';
import axios from 'axios';
import './JsonPrettifier.css';

const JsonPrettifier: React.FC = () => {
  const [inputJson, setInputJson] = useState('');
  const [outputJson, setOutputJson] = useState('');
  const [error, setError] = useState('');
  const [useGraphQL, setUseGraphQL] = useState(false);

  const handleInputChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    setInputJson(e.target.value);
  };

  const formatJsonWithREST = async () => {
    try {
      const response = await axios.post('http://localhost:8080/api/format', {
        json: inputJson
      });
      setOutputJson(response.data.formatted);
      setError('');
    } catch (err) {
      setError('Invalid JSON input');
      setOutputJson('');
    }
  };

  const formatJsonWithGraphQL = async () => {
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

  const handleFormat = () => {
    if (useGraphQL) {
      formatJsonWithGraphQL();
    } else {
      formatJsonWithREST();
    }
  };

  return (
    <div className="json-prettifier">
      <h1>JSON Prettifier</h1>
      
      <div className="api-toggle">
        <label>
          <input
            type="checkbox"
            checked={useGraphQL}
            onChange={() => setUseGraphQL(!useGraphQL)}
          />
          Use GraphQL API
        </label>
      </div>
      
      <div className="input-container">
        <h2>Input JSON</h2>
        <textarea
          value={inputJson}
          onChange={handleInputChange}
          placeholder="Paste your JSON here..."
          rows={10}
        />
      </div>
      
      <button onClick={handleFormat}>Format JSON</button>
      
      {error && <div className="error">{error}</div>}
      
      <div className="output-container">
        <h2>Formatted JSON</h2>
        <pre>{outputJson}</pre>
      </div>
      
      <style jsx>{`
        .json-prettifier {
          max-width: 800px;
          margin: 0 auto;
          padding: 20px;
          font-family: Arial, sans-serif;
        }
        
        h1 {
          text-align: center;
          color: #333;
        }
        
        .input-container, .output-container {
          margin-bottom: 20px;
        }
        
        textarea {
          width: 100%;
          padding: 10px;
          border: 1px solid #ccc;
          border-radius: 4px;
          font-family: monospace;
        }
        
        pre {
          background-color: #f5f5f5;
          padding: 10px;
          border-radius: 4px;
          overflow-x: auto;
          white-space: pre-wrap;
          font-family: monospace;
        }
        
        button {
          background-color: #4CAF50;
          color: white;
          padding: 10px 15px;
          border: none;
          border-radius: 4px;
          cursor: pointer;
          font-size: 16px;
          margin: 10px 0;
        }
        
        button:hover {
          background-color: #45a049;
        }
        
        .error {
          color: red;
          margin: 10px 0;
        }
        
        .api-toggle {
          margin: 10px 0;
        }
      `}</style>
    </div>
  );
};

export default JsonPrettifier; 