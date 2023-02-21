import React, {
  useEffect
} from 'react';
import PropTypes from 'prop-types';

import {
  useMap,
  CircleMarker,
  Popup,
  Polyline
} from 'react-leaflet';

function Markers({ markers }) {
  const map = useMap();

  useEffect(() => {
    map.setView(
      // Center
      [ 33.883136, -117.708721],

      // Zoom
      10
    );
  }, [map]);

  return (
    <>
      {
        // Just use 3 customers because 4 seems too crowded
        markers.slice(0, -1).map(({ id, name, businessCategory, payload, color, crdList }) => (
          crdList.map(({lat, long}, ix) => {
            const radius = ix === crdList.length - 1 ? 7 : 4
            const weight = ix === crdList.length - 1 ? 5 : 6;
            const currentPostion = ix === crdList.length - 1 && 'Current position';

            return (
              <React.Fragment key={id * lat * long}>
                {
                  ix > 0 && (
                    <Polyline
                      dashArray={[5, 5]}
                      positions={[
                        [crdList[ix - 1].lat, crdList[ix - 1].long], [lat, long],
                      ]}
                      weight={1}
                      color={color} />
                  )
                }
                <CircleMarker center={[lat, long]} radius={radius} color={color} weight={weight}>
                  <Popup>
                    {currentPostion && <div>{currentPostion}</div>}
                    {name}
                    <br />
                    Business: {businessCategory}
                    <br />
                    Objects being tracked: {payload}
                  </Popup>
                </CircleMarker>
              </React.Fragment>
            );
          })
        ))
      }
    </>
  );
}

Markers.propTypes = {
  markers: PropTypes.arrayOf(PropTypes.shape()).isRequired
};

export default Markers;
