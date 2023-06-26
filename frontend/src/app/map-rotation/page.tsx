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
        
        <h2>Battle Royal</h2>
        <h3>Current Map</h3>
        <p>{map_rotation.battle_royale.current.map}</p>
        <p>Remain: {map_rotation.battle_royale.current.remainingTimer}</p>
        <Image
          src={map_rotation.battle_royale.current.asset}
          width={600}
          height={250}
          alt={map_rotation.battle_royale.current.map}
        />

        <h3>Next Map</h3>
        <p>{map_rotation.battle_royale.next.map}</p>
        <Image
          src={map_rotation.battle_royale.next.asset}
          width={600}
          height={250}
          alt={map_rotation.battle_royale.next.map}
        />

        <h2>Ranked</h2>
        <h3>Current Map</h3>
        <p>{map_rotation.ranked.current.map}</p>
        <p>Remain: {map_rotation.ranked.current.remainingTimer}</p>
        <Image
          src={map_rotation.ranked.current.asset}
          width={600}
          height={250}
          alt={map_rotation.ranked.current.map}
        />

        <h3>Next Map</h3>
        <p>{map_rotation.ranked.next.map}</p>
        <Image
          src={map_rotation.ranked.next.asset}
          width={600}
          height={250}
          alt={map_rotation.ranked.next.map}
        />
      </nav>
      {children}
    </section>
  )
}

export default MapRotation;