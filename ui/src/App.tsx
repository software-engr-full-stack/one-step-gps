import React from 'react';
import Button from '@mui/material/Button';
import Link from '@mui/material/Link';

import './App.css';

import {useInterval} from './utils';

import Map from './Map/Map';

function App() {
  const [customers, setCustomers] = React.useState([]);
  const [isPollingCoordinates, setIsPollingCoordinates] = React.useState(true);
  const dataDefault : any[] = [];
  const [data, setData] = React.useState(dataDefault);

  React.useEffect(() => {
    fetch('/api/v1/customers/latest').then((json) => json.json()).then((resp) => {
      setCustomers(resp);
    });
  });

  useInterval(() => {
    if (customers.length < 1) return null;
    if (!isPollingCoordinates) return null;

    const ids = JSON.stringify(customers.map(({id}) => id));
    fetch(`/api/v1/customers/coords?ids=${ids}`).then((json) => json.json()).then((resp) => {
      const coordsById = resp.reduce((memo: any, obj: any) => ({
        ...memo,
        [obj.id]: { lat: obj.lat, long: obj.long }
      }), {});

      const list = data.length > 0 ? data : customers;

      const withCoords : any[] = list.map((cust: any) => {
        const currentCrd = coordsById[cust.id];
        if (!currentCrd) throw(`unable to find coordinates for id '${cust.id}'`);

        const initCrdList = cust.crdList ? [...cust.crdList, currentCrd] : [currentCrd];
        const crdList = initCrdList.length > 3 ? initCrdList.slice(cust.crdList.length - 3) : initCrdList;

        return {
          ...cust,
          crdList
        };
      });

      setData(withCoords);
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
      <div className="misc">
        <Button
          onClick={onClick}
          color={isPollingCoordinates ? 'error' : 'success'}
          variant="contained"
        >
          {isPollingCoordinates ? 'Stop' : 'Run'}
        </Button>

        {'    '}

        <Link
          href="https://github.com/software-engr-full-stack/one-step-gps"
          rel="noopener noreferrer"
          target="_blank"
        >
          What is this? (Repo README)
        </Link>
      </div>
      <Map markers={data} />
    </div>
  );
}

export default App;
