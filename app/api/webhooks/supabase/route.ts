// app\api\webhooks\supabase\route.tsx
import { createClient } from '@/utils/supabase/server';
import { NextApiRequest, NextApiResponse } from 'next';
import { Database } from '@/types_db';

export async function GET(req: NextApiRequest, res: NextApiResponse) {
  try {
    const supabase = createClient();
    console.log('Supabase client created.');

    const { data: userData, error: userError } = await supabase.auth.getUser();
    if (userError) {
      console.error('User Error:', userError.message);
      return res.status(500).json({ error: userError.message });
    }
    console.log('User Data:', userData);

    const { data: userDetailsData, error: userDetailsError } = await supabase
      .from('users')
      .select('*')
      .single();
    if (userDetailsError) {
      console.error('User Details Error:', userDetailsError.message);
      return res.status(500).json({ error: userDetailsError.message });
    }
    console.log('User Details Data:', userDetailsData);

    const { data: subscriptionData, error: subscriptionError } = await supabase
      .from('subscriptions')
      .select('*, prices(*, products(*))')
      .in('status', ['trialing', 'active'])
      .maybeSingle();
    if (subscriptionError) {
      console.error('Subscription Error:', subscriptionError.message);
      return res.status(500).json({ error: subscriptionError.message });
    }
    console.log('Subscription Data:', subscriptionData);

    console.log('API Response:', {
      user: userData.user,
      userDetails: userDetailsData,
      subscription: subscriptionData,
    });

    res.status(200).json({
      user: userData.user,
      userDetails: userDetailsData,
      subscription: subscriptionData,
    });
  } catch (error) {
    console.error('Unexpected Error:', error);
    res.status(500).json({ error: 'Unexpected Error' });
  }
}