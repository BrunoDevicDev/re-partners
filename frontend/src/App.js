import React, { useState, useEffect } from 'react';
import './App.css';

function App() {
  const [inputValue, setInputValue] = useState('');
  const [tableData, setTableData] = useState([]);
  const [calculateValue, setCalculateValue] = useState('');
  const [serverResponse, setServerResponse] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  useEffect(() => {
    // Load initial pack sizes from the server
    fetchPackSizes();
  }, []);

  const fetchPackSizes = () => {
    fetch('http://ec2-16-171-20-113.eu-north-1.compute.amazonaws.com:8080/get-parameters')
      .then(response => {
        if (!response.ok) {
          throw new Error('Failed to fetch pack sizes');
        }
        return response.json();
      })
      .then(data => {
        setTableData(data.result); // Ensure correct property is used
      })
      .catch(error => console.error('Error:', error));
  };

  const handleInputChange = (e) => {
    setInputValue(e.target.value);
  };

  const handleAddToTable = () => {
    if (inputValue.trim() !== '') {
      // Add the input value to the end of the table data
      setTableData([...tableData, parseInt(inputValue)]);
      setInputValue(''); // Clear input value after adding
      sendPackSizes([...tableData, parseInt(inputValue)]); // Send updated pack sizes to the server
    } else {
      alert('Please enter a value!');
    }
  };

  const handleChangeValue = (index) => {
    const newValue = prompt('Enter the new value:', tableData[index]);
    if (newValue !== null) {
      const newData = [...tableData];
      newData[index] = parseInt(newValue);
      setTableData(newData);
      sendPackSizes(newData); // Send updated pack sizes to the server
    }
  };

  const handleDeleteValue = (index) => {
    const newData = [...tableData];
    newData.splice(index, 1);
    setTableData(newData);
    sendPackSizes(newData); // Send updated pack sizes to the server
  };

  const sendPackSizes = (packSizes) => {
    // Sending a POST request to the server with the pack sizes
    fetch('http://ec2-16-171-20-113.eu-north-1.compute.amazonaws.com:8080/parameters', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ pack_sizes: packSizes }),
    })
      .then(response => response.json())
      .then(data => {
        console.log(data);
      })
      .catch(error => console.error('Error:', error));
  };

  const handleCalculate = () => {
    if (calculateValue.trim() !== '') {
      setIsLoading(true);
      // Sending a POST request to the server with the number entered
      fetch('http://ec2-16-171-20-113.eu-north-1.compute.amazonaws.com:8080/solve', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ number: parseInt(calculateValue) }),
      })
        .then(response => response.json())
        .then(data => {
          console.log(data);
          if (typeof data.result === 'string') {
            setServerResponse(data.result); // Set server response to state
          } else {
            setServerResponse(JSON.stringify(data.result)); // Convert object to string if needed
          }
          setIsLoading(false);
        })
        .catch(error => {
          console.error('Error:', error);
          setIsLoading(false);
        });
    } else {
      alert('Please enter a value for calculation!');
    }
  };

  return (
    <div className="App">
      <header className="App-header">
        <div className="container">
          <h2>Pack sizes</h2>
          <div className="table">
            <div className="row">
              {tableData.map((item, index) => (
                <div key={index} className="cell">
                  {item}
                  <div className="button-container">
                    <button onClick={() => handleChangeValue(index)}>Change</button>
                    <button onClick={() => handleDeleteValue(index)}>Delete</button>
                  </div>
                </div>
              ))}
            </div>
          </div>
          <div className="black-line"></div>
          <div className="controls">
            <input
              type="text"
              value={inputValue}
              onChange={handleInputChange}
              placeholder="Enter value"
            />
            <button onClick={handleAddToTable}>Add</button>
          </div>
          <div className="black-line"></div>
          <div className="controls">
            <input
              type="text"
              value={calculateValue}
              onChange={(e) => setCalculateValue(e.target.value)}
              placeholder="Enter value for calculation"
            />
            <button onClick={handleCalculate} disabled={isLoading}>
              {isLoading ? 'Calculating...' : 'Calculate'}
            </button>
          </div>
          {serverResponse && <div>Result: {serverResponse}</div>}
        </div>
      </header>
    </div>
  );
}

export default App;
