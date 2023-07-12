'use client'

import { useState, useEffect } from 'react';
import Image from 'next/image'

type RotationInfo = {
  map: string,
  remainingTimer?: string,
  asset: string,
  remainingSecs: number,
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
  const targetSeconds = rotationInfo.remainingSecs;
  const [timeRemaining, setTimeRemaining] = useState(targetSeconds);

  // 毎秒デクリメントする処理
  useEffect(() => {
    const interval = setInterval(() => {
      setTimeRemaining(timeRemaining => timeRemaining - 1);
    }, 1000);

    return () => clearInterval(interval);
  }, []);

  // 秒形式の時刻を時間、分、秒に変換する
  const hoursLeft = Math.floor(timeRemaining / 60 / 60);
  const minutesLeft = Math.floor((timeRemaining / 60) % 60);
  const secondsLeft = timeRemaining % 60;

  return (
    <>
      { isCurrentMap && <h2>{title}</h2> }
      <h3>{isCurrentMap ? "Current Map" : "Next Map"}</h3>
      <p>{rotationInfo.map}</p>
      { isCurrentMap && <p>Remain: {hoursLeft}時間 {minutesLeft}分 {secondsLeft}秒</p> }
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