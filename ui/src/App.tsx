import React from 'react';
import './App.css';

import {useInterval} from './utils';

function App() {
  const [customers, setCustomers] = React.useState([]);
  const [isPollingCoordinates, setIsPollingCoordinates] = React.useState(false);
  const [coordinates, setCoordinates] = React.useState([]);

  React.useEffect(() => {
    fetch('/api/v1/customers/latest').then((resp) => resp.json()).then((data) => {
      setCustomers(data);
    });
  });

  useInterval(() => {
    if (customers.length < 1) return null;
    if (!isPollingCoordinates) return null;

    const ids = JSON.stringify(customers.map(({id}) => id));
    fetch(`/api/v1/customers/coords?ids=${ids}`).then((resp) => resp.json()).then((data) => {
      setCoordinates(data);
    });
  }, 1000 * 2);

  if (customers.length < 1) {
    return null;
  }

  const onClick = () => {
    setIsPollingCoordinates(!isPollingCoordinates);
  };

  return (
    <div className="one-step-gps">
      <button onClick={onClick}>{isPollingCoordinates ? 'Stop' : 'Get coordinates'}</button>
      {
        coordinates.length > 0 && coordinates.map((crd: any) => (
          <React.Fragment key={crd.id}>
            <h3>ID: {crd.id}</h3>
            <h3>Lat: {crd.lat}</h3>
            <h3>Long: {crd.long}</h3>
          </React.Fragment>
        ))
      }
    </div>
  );
}

export default App;
