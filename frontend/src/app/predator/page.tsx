import Header from '../../components/Header';
import Sidebar from '../../components/Sidebar';

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
    <>
      <Header />
      <Sidebar />
      <section>
        <nav>
          <h1>Predator Border</h1>
          
          <h2>PC</h2>
          <p>border: {predator.RP.PC.val}</p>
          <p>P / (M + P): {(750 / predator.RP.PC.totalMastersAndPreds) * 100} %</p>

          <h2>PS4</h2>
          <p>border: {predator.RP.PS4.val}</p>
          <p>P / (M + P): {(750 / predator.RP.PS4.totalMastersAndPreds) * 100} %</p>

          <h2>XBox</h2>
          <p>border: {predator.RP.X1.val}</p>
          <p>P / (M + P): {(750 / predator.RP.X1.totalMastersAndPreds) * 100} %</p>

          <h2>Switch</h2>
          <p>border: {predator.RP.SWITCH.val}</p>
          <p>P / (M + P): {(750 / predator.RP.SWITCH.totalMastersAndPreds) * 100} %</p>
        </nav>
        {children}
      </section>
    </>
  )
}

export default Predator;