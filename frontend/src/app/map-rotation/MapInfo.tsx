import Image from 'next/image'
import MapInfoPanel from './MapInfoPanel'

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
  title,
}: {
  rotationInfo: RotationInfo, 
  title: string
}) => {
  return (
    <>
      <MapInfoPanel
        rotationInfo={rotationInfo.current}
        title={title}
        isCurrentMap={true}
      />
      <MapInfoPanel
        rotationInfo={rotationInfo.next}
        title={title}
        isCurrentMap={false}
      />
    </>
  );
}

export default MapInfo;