import Image from 'next/image'

type RotationInfo = {
  map: string,
  remainingTimer?: string,
  asset: string
}

const MapInfoPanel = ({
  rotationInfo,
  title,
  isCurrentMap
}: {
  rotationInfo: RotationInfo, 
  title: string
  isCurrentMap: boolean
}) => {
  return (
    <>
      <h2>{title}</h2>
      <h3>{isCurrentMap ? "Current Map" : "Next Map"}</h3>
      <p>{rotationInfo.map}</p>
      { isCurrentMap && <p>Remain: {rotationInfo.remainingTimer}</p> }
      <Image
        src={rotationInfo.asset}
        width={600}
        height={250}
        alt={rotationInfo.map}
      />
    </>
  );
}

export default MapInfoPanel;