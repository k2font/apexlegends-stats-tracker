import MapInfo from './MapInfo';
import Header from '../..//components/Header';
import Sidebar from '../../components/Sidebar';
import './styles.css';

const getMapRotation = async () => {
  const res = await fetch("http://localhost:8080/map-rotation")
  return res.json()
}

const MapRotation = async ({
  children,
}: {
  children: React.ReactNode
}) => {
  const map_rotation = await getMapRotation()

  return (
    <div className="map-rotation">
      <div className="map-rotation-header">
        <Header />
      </div>
      <div className="map-rotation-sidebar">
        <Sidebar />
      </div>
      <section>
        <nav>
          <div className="map-rotation-title">
            <h1>Map Rotation</h1>
          </div>
          <div className="container">
            <div className="battle-royale">
              <MapInfo
                rotationInfo={map_rotation.battle_royale}
                title="Battile Royale"
              />
            </div>
            <div className="ranked">
              <MapInfo
                rotationInfo={map_rotation.ranked}
                title="Ranked"
              />
            </div>
          </div>
        </nav>
        {children}
      </section>
    </div>
  )
}

export default MapRotation;