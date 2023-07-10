"use client"

import Header from '../../components/Header';
import Sidebar from '../../components/Sidebar';
import React, {useEffect, useState} from 'react'

export type ServerStatusData = {
  Origin_login: {
    "EU-West": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "EU-East": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-West": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-Central": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-East": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    SouthAmerica: {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    Asia: {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
  }
  EA_novafusion: {
    "EU-West": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "EU-East": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-West": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-Central": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-East": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    SouthAmerica: {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    Asia: {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
  }
  EA_accounts: {
    "EU-West": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "EU-East": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-West": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-Central": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-East": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    SouthAmerica: {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    Asia: {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
  }
  ApexOauth_Crossplay: {
    "EU-West": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "EU-East": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-West": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-Central": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "US-East": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    SouthAmerica: {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    Asia: {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
  }
  selfCoreTest: {
    "Status-website": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "Stats-API": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "Overflow-#1": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "Overflow-#2": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "Origin-API": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "Playstation-API": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
    "Xbox-API": {
      Status: string
      HTTPCode: number
      ResponseTime: number
      QueryTimestamp: number
    }
  }
  otherPlatforms: {
    "Playstation-Network": {
      Status: string
      QueryTimestamp: number
    }
    "Xbox-Live": {
      Status: string
      QueryTimestamp: number
    }
  }
}

const getServerStatus = async () => {
  const res = await fetch("http://localhost:8080/server-status")
  return res.json()
}

const ServerStatus = ({
  children,
}: {
  children: React.ReactNode
}) => {
  const [status, setStatus] = useState<ServerStatusData | null>(null)

  useEffect(() => {
    const fetchData = async () => {
      try {
        const statusData = await getServerStatus();
        setStatus(statusData);
        console.log("test");
      } catch (error) {
        console.error('Error fetching server status:', error);
      }
    };
    fetchData();
  }, []);

  return (
    <>
      <Header />
      <Sidebar />
      {!status ? (
        <p>Loading...</p>
      ) : (
        <section>
          <nav>
            <h1>Server Status</h1>

            <h2>Origin Login</h2>
            <p>EUWest: {status.Origin_login["EU-West"].Status}</p>
            <p>EUeast: {status.Origin_login["EU-East"].Status}</p>
            <p>USWest: {status.Origin_login["US-West"].Status}</p>
            <p>USCentral: {status.Origin_login["US-Central"].Status}</p>
            <p>USEast: {status.Origin_login["US-East"].Status}</p>
            <p>SouthAmerica: {status.Origin_login["SouthAmerica"].Status}</p>
            <p>Asia: {status.Origin_login["Asia"].Status}</p>
          </nav>
          {children}
        </section>
      )}
    </>
  )
}

export default ServerStatus;