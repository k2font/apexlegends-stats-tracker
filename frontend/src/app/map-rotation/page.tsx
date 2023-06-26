import MapInfo from './MapInfo';

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
        <h1>Map Rotation</h1>
        <MapInfo
          rotationInfo={map_rotation.battle_royale}
          title="Battile Royale"
        />
        <MapInfo
          rotationInfo={map_rotation.ranked}
          title="Ranked"
        />
      </nav>
      {children}
    </section>
  )
}

export default MapRotation;