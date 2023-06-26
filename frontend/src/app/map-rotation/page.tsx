import Image from 'next/image'

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
        
        <h2>Current Map</h2>
        <p>{map_rotation.current.map}</p>
        <p>Remain: {map_rotation.current.remainingTimer}</p>
        <Image
          src={map_rotation.current.asset}
          width={1200}
          height={600}
          alt={map_rotation.current.map}
        />
      </nav>
      {children}
    </section>
  )
}

export default MapRotation;