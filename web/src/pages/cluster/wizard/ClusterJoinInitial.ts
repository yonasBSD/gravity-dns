import { ClusterInstancesApi, RolesApiApi } from "gravity-api";

import { customElement } from "@lit/reactive-element/decorators/custom-element.js";
import { TemplateResult, html } from "lit";

import { DEFAULT_CONFIG } from "../../../api/Config";
import { KeyUnknown } from "../../../elements/forms/Form";
import "../../../elements/forms/FormGroup";
import "../../../elements/forms/HorizontalFormElement";
import { WizardFormPage } from "../../../elements/wizard/WizardFormPage";
import { Roles } from "../RolesPage";

export const joinUserUsername = "gravity-system:cluster:join";

@customElement("gravity-cluster-join-initial")
export class ClusterJoinInitial extends WizardFormPage {
    sidebarLabel = () => "Node details";

    nextDataCallback = async (data: KeyUnknown): Promise<boolean> => {
        this.host.state["identifier"] = data.name;

        const roles: string[] = [];
        const prefix = "role_";
        Object.keys(data).forEach((key) => {
            if (!key.startsWith(prefix)) {
                return;
            }
            if (!data[key]) {
                return;
            }
            roles.push(key.replaceAll(prefix, ""));
        });
        this.host.state["roles"] = roles.join(";");

        const info = await new ClusterInstancesApi(DEFAULT_CONFIG).clusterGetInstanceInfo();
        this.host.state["node_ip"] = info.instanceIP;

        await new RolesApiApi(DEFAULT_CONFIG).apiPutUsers({
            username: joinUserUsername,
            authAPIUsersPutInput: {
                password: "",
                permissions: [
                    {
                        methods: ["POST"],
                        path: "/api/v1/etcd/join",
                    },
                ],
            },
        });

        const token = await new RolesApiApi(DEFAULT_CONFIG).apiPutTokens({
            username: joinUserUsername,
        });
        this.host.state["join_token"] = token.key;
        return true;
    };

    renderForm(): TemplateResult {
        return html`<ak-form-element-horizontal label=${"Name"} required name="name">
                <input type="text" value="" class="pf-c-form-control" required />
                <p class="pf-c-form__helper-text">
                    The unique identifier of the node being added to the cluster.
                </p>
            </ak-form-element-horizontal>
            <ak-form-element-horizontal label=${"Roles"} required>
                ${Roles.map((role) => {
                    return html`<div class="pf-c-check">
                        <input
                            type="checkbox"
                            class="pf-c-check__input"
                            checked
                            name=${`role_${role.id}`}
                        />
                        <label class="pf-c-check__label"> ${role.name} </label>
                    </div>`;
                })}
                <p class="pf-c-form__helper-text">Select which roles the new node should provide</p>
            </ak-form-element-horizontal>`;
    }
}
