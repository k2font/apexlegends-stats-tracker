const getServerStatus = async () => {
  const res = await fetch("http://localhost:8080/server-status")
  return res.json()
}

const ServerStatus = async ({
  children,
}: {
  children: React.ReactNode
}) => {
  const status = await getServerStatus()

  return (
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
  )
}

export default ServerStatus;