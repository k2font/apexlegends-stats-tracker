import Image from 'next/image'

type RotationInfo = {
  current: {
    map: string,
    remainingTimer: string,
    asset: string
  },
  next: {
    map: string,
    asset: string
  },
}

const MapInfo = ({
  rotationInfo,
  title
}: {
  rotationInfo: RotationInfo, 
  title: string
}) => {
  return (
    <>
      <h2>{title}</h2>
      <h3>Current Map</h3>
      <p>{rotationInfo.current.map}</p>
      <p>Remain: {rotationInfo.current.remainingTimer}</p>
      <Image
        src={rotationInfo.current.asset}
        width={600}
        height={250}
        alt={rotationInfo.current.map}
      />

      <h3>Next Map</h3>
      <p>{rotationInfo.next.map}</p>
      <Image
        src={rotationInfo.next.asset}
        width={600}
        height={250}
        alt={rotationInfo.next.map}
      />
    </>
  );
}

export default MapInfo;