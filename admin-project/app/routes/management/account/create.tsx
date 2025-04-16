import { FormLayout, FormInput, FormDropDown } from "~/components/Form";
import { SaveBtn } from "~/components/Form/SaveBtn";
import BackButton from "~/components/common/BackButton";
import { useNavigate } from "react-router";
import { ManagerAPI } from "./endpoint";

export default function AccountCreatePage() {
    const navigate = useNavigate();

    return (
        <div className="w-[400px] ml-4 mt-8">
            <div className="flex items-center justify-between mb-6">
                <h2 className="text-xl font-bold">Create</h2>
                <BackButton label="<-" />
            </div>

            <FormLayout>
                <FormInput name="id" placeholder="id" />
                <FormInput name="name" placeholder="user name" />
                <FormDropDown tab= "manager.grade.type" name="grade"/>      
                <SaveBtn
                    endpoint={ManagerAPI.create}
                    onSuccess={() => navigate("..")}
                />
            </FormLayout>
        </div>
    );
}
