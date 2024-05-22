import CustomerPortalForm from '@/components/ui/AccountForms/CustomerPortalForm';
import EmailForm from '@/components/ui/AccountForms/EmailForm';
import NameForm from '@/components/ui/AccountForms/NameForm';
import { createClient } from '@/utils/supabase/server';
import { redirect } from 'next/navigation';
import { powerbiApi } from '@/powerbi.api';
import { powerbiConfig } from '@/powerbi.config';
import { useEffect, useState } from 'react';

export default function SiteScoutDash() {
  const [dashboardUrl, setDashboardUrl] = useState<string | null>(null); // Correctly typed useState
  const [userDetails, setUserDetails] = useState<any>(null);
  const [subscription, setSubscription] = useState<any>(null);
  const [user, setUser] = useState<any>(null);
  const supabase = createClient();

  useEffect(() => {
    async function fetchDashboard() {
      const { data: userData } = await supabase.auth.getUser();
      setUser(userData.user);

      if (!userData.user) {
        redirect('/signin');
        return;
      }

      const { data: userDetailsData } = await supabase
        .from('users')
        .select('*')
        .single();
      setUserDetails(userDetailsData);

      const { data: subscriptionData, error } = await supabase
        .from('subscriptions')
        .select('*, prices(*, products(*))')
        .in('status', ['trialing', 'active'])
        .maybeSingle();
      setSubscription(subscriptionData);

      if (error) {
        console.log(error);
      }

      const accessToken = await powerbiApi.authenticate();
      const dashboard = await powerbiApi.getDashboard(accessToken);
      if (dashboard) {
        setDashboardUrl(`https://app.powerbi.com/reportEmbed?groupId=${powerbiConfig.groupId}&dashboardId=${powerbiConfig.dashboardId}`);
      }
    }

    fetchDashboard();
  }, []);

  return (
    <section className="mb-32 bg-black">
      <div className="max-w-6xl px-4 py-8 mx-auto sm:px-6 sm:pt-24 lg:px-8">
        <div className="sm:align-center sm:flex sm:flex-col">
          <h1 className="text-4xl font-extrabold text-white sm:text-center sm:text-6xl">
            Site Scout Dashboard
          </h1>
          <p className="max-w-2xl m-auto mt-5 text-xl text-zinc-200 sm:text-center sm:text-2xl">
            We partnered with Stripe for a simplified billing.
          </p>
        </div>
      </div>
      <div className="p-4">
        {subscription && <CustomerPortalForm subscription={subscription} />}
        {userDetails && <NameForm userName={userDetails.full_name ?? ''} />}
        {user && <EmailForm userEmail={user.email} />}
        {dashboardUrl && (
          <iframe
            src={dashboardUrl}
            frameBorder="0"
            width="100%"
            height="800"
            allowFullScreen={true}
          />
        )}
      </div>
    </section>
  );
}