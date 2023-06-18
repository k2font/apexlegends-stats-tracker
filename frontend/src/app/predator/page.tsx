const getPredator = async () => {
  const res = await fetch("http://localhost:8080/predator")
  return res.json()
}

const Predator = async ({
  children,
}: {
  children: React.ReactNode
}) => {
  const predator = await getPredator()
  console.log(predator)

  return (
    <section>
      <nav>
        <h1>Predator Border</h1>
        
        <p>PC: {predator.RP.PC.totalMastersAndPreds}</p>
        <p>PS4/PS5: {predator.RP.PS4.totalMastersAndPreds}</p>
        <p>XBox: {predator.RP.X1.totalMastersAndPreds}</p>
        <p>Switch: {predator.RP.SWITCH.totalMastersAndPreds}</p>
      </nav>
      {children}
    </section>
  )
}

export default Predator;