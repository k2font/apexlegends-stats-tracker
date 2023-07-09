import MapInfo from './MapInfo';
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
    <section>
      <nav>
        <div>
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
  )
}

export default MapRotation;