import Express, { Request, Response } from 'express'
import dotenv from 'dotenv'
import Stripe from 'stripe'
dotenv.config()

const stripe = new Stripe(process.env.STRIPE_SECRET_KEY!, {
  apiVersion: '2022-08-01',
  typescript: true,
})

const app = Express()
app.use(Express.json())

app.post('/payment/create', async (req: Request, res: Response) => {
  const { amount, userId, description } = req.body
  try {
    const paymentIntent = await stripe.paymentIntents.create({
      amount: amount * 100,
      currency: 'INR',
      customer: userId,
      payment_method: 'pm_card_in',
      confirm: true,
      description: description,
    })

    /* Add the payment intent record to your datbase if required */
    // Returning the whole paymentIntent object
    // Obtain ID: paymentIntent.id
    res.status(200).json(paymentIntent)
  } catch (err) {
    console.log(err)
    res.status(500).json('Could not create payment')
  }
})

app.post('/payment/confirm', async (req: Request, res: Response) => {
  // Just send the id as paymentIntent
  const { paymentIntent } = req.body
  try {
    const intent = await stripe.paymentIntents.confirm(paymentIntent, {
      payment_method: 'pm_card_in',
    })
    /* Update the status of the payment to indicate confirmation */
    res.status(200).json(intent)
  } catch (err) {
    console.error(err)
    res.status(500).json('Could not confirm payment')
  }
})

app.listen(4000, () => console.log('Server running on port 4000.'))
